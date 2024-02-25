package main

import (
	"os"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/g_var"
	"agent_redis/pkg/worker"
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
	wBuilder := worker.NewWorkerManagerBuilder()

	g_var.InitGlobalVar(agent_id)
	initLogger(cfg)
	initAgentDbClient(cfg)
	initWorkerManagerBuilder(cfg, wBuilder)

	wm := wBuilder.Build()
	wm.StartAndBlock()
}