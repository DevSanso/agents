package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/g_var"
	"agent_redis/pkg/global/log"
)

func main() {

	if(len(os.Args) < 3) {
		panic("Usage: ./execute configPath agent_id")
	}

	configPath := os.Args[1]
	agent_id := os.Args[2]

	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		panic("ReadConfig File Erorr : " + err.Error())
	}

	initGoRuntime()
	
	g_var.InitGlobalVar(agent_id)
	err = initLogger(cfg)
	if err != nil {
		panic("InitLogger Erorr : " + err.Error())
	}
	err = initAgentDbClient(cfg)
	if err != nil {
		panic("InitAgentDBClient Erorr : " + err.Error())
	}

	ctls, workerErr := initWorkers(cfg)

	if workerErr != nil {
		panic("Init Workers Error : " + workerErr.Error())
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	mainIsRun := true

	fmt.Println("Agent is running")

	for mainIsRun {
		select {
			case sig := <-sigs:
				log.GetLogger().Info(fmt.Sprintf("Agent %s(%s) is stopping", agent_id, sig.String()))
				mainIsRun = false
				continue
			default:
				time.Sleep(time.Second * 1)
		}
	}

	for _, ctl := range ctls {
		ctl.Close()
	}
}