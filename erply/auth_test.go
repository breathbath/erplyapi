package erply

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type cacheMock struct {
	StoreKey string
	StoreSource interface{}
	StoreExpiration time.Duration
	StoreErr error

	ReadKey string
	ReadTarget interface{}
	ReadTargetToGive []byte
	ReadFound bool
	ReadErr error

	DeleteKey string
	DeleteErr error
}

//Store interface implementation of Cache
func (cm *cacheMock) Store (key string, source interface{}, expiration time.Duration) error {
	cm.StoreKey = key
	cm.StoreSource = source
	cm.StoreExpiration = expiration

	return cm.StoreErr
}

//Read interface implementation of Cache
func (cm *cacheMock) Read (key string, target interface{}) (found bool, err error) {
	cm.ReadKey = key
	cm.ReadTarget = target

	err = json.Unmarshal(cm.ReadTargetToGive, target)
	if err != nil {
		return false, err
	}

	return cm.ReadFound, cm.ReadErr
}

//Delete interface implementation of Cache
func (cm *cacheMock) Delete (key string) (err error) {
	cm.DeleteKey = key
	return cm.DeleteErr
}

func setEnvs(envMap map[string]string) error {
	for key, val :=  range envMap {
		err := os.Setenv(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestGetSessionFromAPI(t *testing.T) {
	err := setEnvs(map[string] string {
		"ERPLY_CLIENT_CODE": "111",
		"ERPLY_USERNAME": "nn",
		"ERPLY_PASS": "ppp",
	})
	assert.NoError(t, err)
	if err != nil {
		return
	}


	c := &cacheMock{
		ReadErr: nil,
		ReadFound: false,
		ReadTargetToGive: []byte(`""`),
	}

	var un, p, code string
	ap := AuthProvider{
		CacheClient: c,
		AuthFunc: func(username string, password string, clientCode string) (s string, err error) {
			un = username
			p = password
			code = clientCode

			return "sess123", nil
		},
	}

	sess, err := ap.GetSession()
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, "sess123", sess.sessionID)
	assert.Equal(t, "nn", un)
	assert.Equal(t, "ppp", p)
	assert.Equal(t, "111", code)
}

func TestGetSessionFromCache(t *testing.T) {
	err := setEnvs(map[string] string {
		"ERPLY_CLIENT_CODE": "2222",
		"ERPLY_USERNAME": "nn222",
		"ERPLY_PASS": "ppp222",
	})
	assert.NoError(t, err)
	if err != nil {
		return
	}


	c := &cacheMock{
		ReadErr: nil,
		ReadFound: true,
		ReadTargetToGive: []byte(`"sessFromCache111"`),
	}

	ap := AuthProvider{
		CacheClient: c,
		AuthFunc: func(username string, password string, clientCode string) (s string, err error) {
			assert.Fail(t, "Auth func should not be called")
			return
		},
	}

	sess, err := ap.GetSession()
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, "sessFromCache111", sess.sessionID)
}
