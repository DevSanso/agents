package db_test

import (
	"encoding/json"
	"os"
	"testing"
	"context"

	"agent_redis/pkg/db"
)

var commander db.IRedisCoreClientCommander

func init() {
	bytes, err := os.ReadFile("../../config/test/redis.json")
	if err != nil {
		panic(err)
	}
	opts := &db.ClientOptions{}

	err = json.Unmarshal(bytes, opts)
	if err != nil {
		panic(err)
	}
	commander = db.NewCoreClient(opts)
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
	
	t.Logf("%d\n", info)
}

func TestRedisMemoryInfo(t *testing.T) {

	info, infoErr := commander.InfoMemory(context.Background())
	if infoErr != nil {
		t.Error(infoErr)
		return
	}
	
	t.Logf("%+v\n", info)
}
