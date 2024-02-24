package db_test

import (
	"testing"
	"context"

	"agent_redis/pkg/db"
	"agent_redis/pkg/config"
)

var commander db.IRedisCoreClientCommander

func init() {
	cfg, err := config.ReadConfigFromFile("../../assets/config/agent_redis.toml")
	if err != nil {
		panic(err)
	}
	commander = db.NewCoreClient(&db.ClientOptions{
		Timeout: 1,
		Ip : cfg.Redis.Ip,
		Port: cfg.Redis.Port,
		Username: cfg.Redis.UserName,
		Password: cfg.Redis.Password,
		Db: cfg.Redis.Dbname,
		DbVersion: cfg.Redis.DbVersion,
	})
	
}

func TestRedisClientList(t *testing.T) {

	list, infoErr := commander.GetClientList(context.Background())
	if infoErr != nil {
		t.Error(infoErr)
		return
	}
	
	t.Logf("%+v\n", list)
}

func TestRedisInfoStats(t *testing.T) {

	info, infoErr := commander.InfoStat(context.Background())
	if infoErr != nil {
		t.Error(infoErr)
		return
	}
	
	t.Logf("%+v\n", info)
}

func TestRedisInfoCpu(t *testing.T) {

	info, infoErr := commander.InfoCpu(context.Background())
	if infoErr != nil {
		t.Error(infoErr)
		return
	}
	
	t.Logf("%+v\n", info)
}

func TestRedisDbSize(t *testing.T) {

	info, infoErr := commander.GetDbSize(context.Background(),10)
	if infoErr != nil {
		t.Error(infoErr)
		return
	}
	
	t.Logf("%+v\n", info)
}

func TestRedisMemoryInfo(t *testing.T) {

	info, infoErr := commander.InfoMemory(context.Background())
	if infoErr != nil {
		t.Error(infoErr)
		return
	}
	
	t.Logf("%+v\n", info)
}
