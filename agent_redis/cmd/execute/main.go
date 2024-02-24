package main

import (
	"os"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/log"
	g_db "agent_redis/pkg/global/db"
	"agent_redis/pkg/global/g_var"
	"agent_redis/pkg/db"
	"agent_redis/pkg/worker"
)

func main() {

	if(len(os.Args) < 3) {
		panic("Usage: ./execute configPath agent_id")
		return;
	}

	configPath := os.Args[1]
	agent_id := os.Args[2]
	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		panic("ReadConfig File Erorr : " + err.Error())
	}
	err = log.InitLogger(cfg.LogFilePath, log.LogLevel(cfg.LogLevel))
	if err != nil {
		panic("Init Logger Error : " + err.Error())
	}
	err = g_db.InitRedis(&db.ClientOptions{
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
	g_var.InitGlobalVar(agent_id)

	var cfgFuncCaller func() (string,int,error) = nil
	if cfg.Sender.SendType == "TCP" {
		cfgFuncCaller = cfg.Sender.TcpConfig
	}else {
		cfgFuncCaller = cfg.Sender.MmapConfig
	}

	strval,intval, cfgErr := cfgFuncCaller()
	if cfgErr != nil {
		log.GetLogger().Error(cfgErr.Error() + " Type(" + cfg.Sender.SendType+")")
		os.Exit(2)
	}

	var sendWorker worker.IWorker = nil
	var sendWorkerErr error = nil

	if cfg.Sender.SendType == "TCP" {
		//sendWorker,sendWorkerErr = NewTcpSendWorker(strval, intval)
		log.GetLogger().Error("not support TCP Snap IPC")
		os.Exit(2)
	}else {
		sendWorker,sendWorkerErr = NewMmapSendWorker(strval, intval)
	}

	if sendWorkerErr != nil {
		log.GetLogger().Error(sendWorkerErr.Error())
		os.Exit(2)
	}
	
	middleWareWorker := NewMiddleWareWorker()
	clientInfoCmdWorker := NewClientInfoWorker()

	wBuilder := worker.NewWorkerManagerBuilder()
	wBuilder.SendWorker(sendWorker)
	wBuilder.MiddleWareWorker(middleWareWorker)
	wBuilder.AddCmdWorker(clientInfoCmdWorker, time.Second * 3)
	
	wm := wBuilder.Build()
	wm.StartAndBlock()
}