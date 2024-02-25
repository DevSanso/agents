package main

import (
	"os"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/log"
	g_db "agent_redis/pkg/global/db"
	"agent_redis/pkg/db"
	"agent_redis/pkg/worker"
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

func initWorkerManagerBuilder(cfg *config.Config, builder *worker.WorkerManagerBuilder) {
	strval,intval, cfgErr := cfg.Sender.MmapConfig()
	if cfgErr != nil {
		log.GetLogger().Error(cfgErr.Error() + " Type(" + cfg.Sender.SendType+")")
		os.Exit(2)
	}

	var sendWorker worker.IWorker = nil
	var sendWorkerErr error = nil

	sendWorker,sendWorkerErr = NewMmapSendWorker(strval, intval)

	if sendWorkerErr != nil {
		log.GetLogger().Error(sendWorkerErr.Error())
		os.Exit(2)
	}
	
	middleWareWorker := NewMiddleWareWorker()
	clientInfoCmdWorker := NewClientInfoWorker()

	builder.SendWorker(sendWorker)
	builder.MiddleWareWorker(middleWareWorker)
	builder.AddCmdWorker(clientInfoCmdWorker, time.Second * 3)
}