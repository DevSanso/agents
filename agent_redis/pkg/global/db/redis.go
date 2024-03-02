package db

import (
	redis_db "agent_redis/pkg/db"
	"io"
	"sync"
)

var (
	once sync.Once
	core redis_db.IRedisCoreClientCommander
	dbCloser io.Closer
)


func InitRedis(opt *redis_db.ClientOptions) (err error) {
	once.Do(func() {
		core, dbCloser = redis_db.NewCoreClient(opt)
	})

	return
}

func GetCoreClient() redis_db.IRedisCoreClientCommander {
	return core
}

func GetCoreClientCloser() io.Closer {
	return dbCloser
}

