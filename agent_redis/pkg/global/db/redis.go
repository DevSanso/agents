package db

import (
	"sync"
	redis_db "agent_redis/pkg/db"
)

var (
	once sync.Once
	core redis_db.IRedisCoreClientCommander
)


func InitRedis(opt *redis_db.ClientOptions) (err error) {
	once.Do(func() {
		core = redis_db.NewCoreClient(opt)
	})

	return
}

func GetCoreClient() redis_db.IRedisCoreClientCommander {
	return core
}

