package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClientCommander struct {
	client    *redis.Client
	dbVersion float32
}

type IRedisCoreClientCommander interface {
	GetClientList(ctx context.Context) ([]RedisClientInfo, error)
	InfoStat(ctx context.Context) (*RedisStatsInfo, error)
	GetDbSize(ctx context.Context, dbname int) (int64, error)
	InfoCpu(ctx context.Context) (*RedisCpuInfo, error)
}

type IRedisStackClientCommander interface {
	IRedisCoreClientCommander
}

type ClientOptions struct {
	Timeout   int `json:"timeout"`
	Ip        string `json:"ip"`
	Port      int `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Db        int `json:"db"`
	DbVersion float32 `json:"dbversion"`
}

func newRedisClientCommander(timeout int, ip string, port int, username string, passwd string, dbname int, dbVersion float32) *redisClientCommander {

	opts := &redis.Options{
		Addr:                  fmt.Sprintf("%s:%d", ip, port),
		Password:              passwd,
		Username:              username,
		DB:                    dbname,
		MaxRetries:            5,
		MinRetryBackoff:       time.Millisecond * 10,
		MaxRetryBackoff:       time.Millisecond * 100,
		DialTimeout:           time.Second * 1,
		ReadTimeout:           time.Second * 10,
		ContextTimeoutEnabled: true,
		PoolFIFO:              false,
		PoolSize:              10,
		PoolTimeout:           time.Second * 5,
		MaxIdleConns:          10,
		MaxActiveConns:        5,
		ConnMaxIdleTime:       time.Second * 70,
		ConnMaxLifetime:       0,
	}

	c := redis.NewClient(opts)
	return &redisClientCommander{c, dbVersion}
}

func (rcc *redisClientCommander) Close() error {
	return rcc.client.Close()
}

func NewCoreClient(opts *ClientOptions) IRedisCoreClientCommander {
	return newRedisClientCommander(opts.Timeout, opts.Ip, opts.Port, opts.Username, opts.Password, opts.Db, opts.DbVersion)
}
