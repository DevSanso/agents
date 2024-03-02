package main

import (
	"context"
	"time"
	"runtime"

	"agent_redis/cmd/execute/workers"
	"agent_redis/pkg/config"
	"agent_redis/pkg/db"
	g_db "agent_redis/pkg/global/db"
	"agent_redis/pkg/global/log"
	"agent_redis/pkg/worker"
)

func initGoRuntime() error {
	runtime.GOMAXPROCS(12)
	return nil
}

func initLogger(cfg *config.Config) error {
	err := log.InitLogger(cfg.LogFilePath, log.LogLevel(cfg.LogLevel))
	if err != nil {
		return err
	}
	if(len(cfg.LogFilePath) <= 0) {
		log.GetLogger().Warn("LogFilePath is empty, only print to console")
	}
	return nil
}

func initAgentDbClient(cfg *config.Config) error {
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
		return err
	}
	pingCtx, cancleFunc := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancleFunc()
	err = g_db.GetCoreClient().Ping(pingCtx)

	if err != nil {
		log.GetLogger().Error(err.Error())
		return err
	}
	
	return nil
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
		WorkerTimeOut: time.Millisecond * 50,
		WorkerInterval: time.Second * 1,
		SendChan: miidleChannel,
		RecvChan: nil,
	}))

	ret = append(ret, worker.NewWorkerThread(workers.NewStatInfoWorker(), worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Millisecond * 100,
		WorkerInterval: time.Second * 60,
		SendChan: miidleChannel,
		RecvChan: nil,
	}))

	ret = append(ret, worker.NewWorkerThread(workers.NewMemInfoWorker(), worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Millisecond * 100,
		WorkerInterval: time.Second * 3,
		SendChan: miidleChannel,
		RecvChan: nil,
	}))
	ret = append(ret, worker.NewWorkerThread(workers.NewCpuInfoWorker(), worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Millisecond * 100,
		WorkerInterval: time.Second * 3,
		SendChan: miidleChannel,
		RecvChan: nil,
	}))
	
	ret = append(ret, worker.NewWorkerThread(mmapWorker, worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Millisecond * 100,
		WorkerInterval: time.Millisecond * 100,
		SendChan: nil,
		RecvChan: sendChannel,
		WorkerCloser: mmapWorker,
	}))
	ret = append(ret, worker.NewWorkerThread(workers.NewMiddleWareWorker(), worker.WorkerThreadStartUpArgs{
		WorkerTimeOut: time.Millisecond * 90,
		WorkerInterval: time.Millisecond * 90,
		SendChan: sendChannel,
		RecvChan: miidleChannel,
	}))

	return ret, nil
}
