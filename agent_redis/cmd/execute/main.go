package main

import (
	"os"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/log"
	"agent_redis/pkg/worker"
)

func main() {
	configPath := os.Args[1]
	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		panic(err)
	}
	err = log.InitLogger(cfg.LogFilePath)
	if err != nil {
		panic(err)
	}

	var cfgFuncCaller func() (string,int,error) = nil
	if cfg.Sender.SendType == "TCP" {
		cfgFuncCaller = cfg.Sender.TcpConfig
	}else {
		cfgFuncCaller = cfg.Sender.MmapConfig
	}

	strval,intval, cfgErr := cfgFuncCaller()
	if cfgErr != nil {
		log.GetLogger().Error(cfgErr.Error())
		os.Exit(2)
	}

	var sendWorker worker.IWorker = nil
	var sendWorkerErr error = nil

	if cfg.Sender.SendType == "TCP" {
		sendWorker,sendWorkerErr = NewTcpSendWorker(strval, intval)
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