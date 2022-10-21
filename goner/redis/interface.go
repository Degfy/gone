package redis

import "time"

type Cache interface {
	Put(key string, value interface{}, ttl ...time.Duration) error
	Get(key string, value interface{}) error
	Remove(key string) (err error)
	Keys(key string) ([]string, error)
}

type Locker interface {
	TryLock(key string, ttl time.Duration) (unlock Unlock, err error)
	LockAndDo(key string, fn func(), lockTime, checkPeriod time.Duration) (err error)
}