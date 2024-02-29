package main

import (
	"os"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/log"
	g_db "agent_redis/pkg/global/db"
	"agent_redis/pkg/db"
	"agent_redis/pkg/worker"
	"agent_redis/cmd/execute/workers"
)

func initLogger(cfg *config.Config) {
	err := log.InitLogger(cfg.LogFilePath, log.LogLevel(cfg.LogLevel))
	if err != nil {
		panic("Init Logger Error : " + err.Error())
	}
	if(len(cfg.LogFilePath) <= 0) {
		log.GetLogger().Warn("LogFilePath is empty, only print to console")
	}
}

func initAgentDbClient(cfg *config.Config) {
	err := g_db.InitRedis(&db.ClientOptions{
		Timeout: 1,
		Ip : cfg.Redis.Ip,
		Port: cfg.Redis.Port,
		Username: cfg.Redis.UserName,
		Password: cfg.Redis.Password,
		Db: cfg.Redis.Dbname,
		DbVersion: cfg.Redis.DbVersion,
	})
	if err != nil {
		log.GetLogger().Error(err.Error())
		os.Exit(2)
	}
}

func initWorkers(cfg *config.Config) ([]*worker.WorkerThreadCtl, error) {
	filename, size, err := cfg.Sender.MmapConfig()
	if err != nil {
		return nil, err
	}
	mmapWorker,mmapErr := workers.NewMmapSendWorker(filename, size)
	if mmapErr != nil {
		return nil, mmapErr
	
	}

	miidleChannel := make(chan *worker.WorkerResponse)
	sendChannel := make(chan *worker.WorkerResponse)

	ret := make([]*worker.WorkerThreadCtl, 0)

	ret = append(ret, worker.NewWorkerThread(workers.NewClientInfoWorker(), worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Second * 1,
		WorkerInterval: time.Second * 1,
		SendChan: nil,
		RecvChan: miidleChannel,
	}))
	ret = append(ret, worker.NewWorkerThread(mmapWorker, worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Second * 3,
		WorkerInterval: time.Second * 2,
		SendChan: sendChannel,
		RecvChan: nil,
	}))
	ret = append(ret, worker.NewWorkerThread(workers.NewMiddleWareWorker(), worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Second * 1,
		WorkerInterval: time.Second * 2,
		SendChan: miidleChannel,
		RecvChan: sendChannel,
	}))

	return ret, nil
}
