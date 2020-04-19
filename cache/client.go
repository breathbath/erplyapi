package cache

import "time"

//Client interface to abstract concrete cache backend communication
type Client interface {
	Store (key string, source interface{}, expiration time.Duration) error
	Read (key string, target interface{}) (found bool, err error)
	Delete (key string) (err error)
}
