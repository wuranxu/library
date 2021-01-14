package redis

import (
	"fmt"
	rds "github.com/go-redis/redis"
	"sync"
)

var (
	Pool *rds.Client
	once sync.Once
)

func NewClient(ops *rds.Options) (*rds.Client, error) {
	var err error
	if Pool == nil {
		once.Do(func() {
			Pool = rds.NewClient(ops)
			if _, err = Pool.Ping().Result(); err != nil {
				err = fmt.Errorf("connect redis failed, error: %v", err)
			}
		})
	}
	if err != nil {
		return nil, err
	}
	return Pool, nil
}
