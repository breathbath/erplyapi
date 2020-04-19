package redis

import (
	"encoding/json"
	"github.com/breathbath/erplyapi/utils"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

//Client redis implementation of cache.Client
type Client struct {
	MarshalFunc func (v interface{}) ([]byte, error)
	UnMarshalFunc func (payload []byte, target interface{}) error
}

//Store redis implementation of a corresponding method in cache.Client
func (c Client) Store(key string, source interface{}, expiration time.Duration) error {
	bcl := c.newBaseClient()
	defer utils.Close(bcl)

	payload, err := c.marshalSource(source)
	if err != nil {
		return err
	}

	log.Debugf("Will store '%s' in Redis under key '%s' with expiration '%v'", key, string(payload), expiration)

	err = bcl.Set(key, payload, expiration).Err()
	if err != nil {
		return err
	}

	log.Debug("Successfully stored")

	return nil
}

//Read redis implementation of a corresponding method in cache.Client
func (c Client) Read(key string, target interface{}) (found bool, err error) {
	bcl := c.newBaseClient()
	defer utils.Close(bcl)

	log.Debugf("Will read '%s' from Redis", key)

	payload, err := bcl.Get(key).Bytes()

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		return false, err
	}

	log.Debugf("Successful read data: '%s' from Redis", string(payload))

	err = c.unmarshalToTarget(payload, target)
	if err != nil {
		return
	}

	return true, nil
}

//Delete deletes a key
func (c Client) Delete (key string) (err error) {
	log.Debugf("Will delete key '%s' from Redis", key)
	bcl := c.newBaseClient()
	defer utils.Close(bcl)

	err = bcl.Del(key).Err()
	if err == nil {
		log.Debugf("Successfully deleted")
	}
	return
}

func (c Client) newBaseClient() *redis.Client {
	redisConnStr := env.ReadEnv("REDIS_CONN_STR", "localhost:6379")
	client := redis.NewClient(&redis.Options{
		Addr:     redisConnStr,
		Password: env.ReadEnv("REDIS_PASS", ""),
		DB:       0,
	})
	return client
}

func (c Client) marshalSource(source interface{}) (payload []byte, err error) {
	if c.MarshalFunc == nil {
		payload, err = json.Marshal(source)
		return
	}

	payload, err = c.MarshalFunc(source)
	return
}

func (c Client) unmarshalToTarget(payload []byte, target interface{}) (err error) {
	if c.UnMarshalFunc == nil {
		err = json.Unmarshal(payload, target)
		return
	}

	err = c.UnMarshalFunc(payload, target)
	return
}
