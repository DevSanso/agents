package main

import (
	"os"
	"fmt"
	"database/sql"

	"sched/pkg/conf"
	"sched/pkg/global"
)

func InitDbConn(c *conf.Configure) (*sql.DB,error) {
	dbConf := c.DBConfig;
	return sql.Open("postgres", 
		fmt.Sprintf("user=%s password='%s' host=%s port=%d dbname=%s", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName))
}

func main() {
	confPath := os.Args[1]
	globalErr := global.NewGlobal(confPath)
	
	if globalErr != nil {
		panic(globalErr)
	}

	db, dbErr := InitDbConn(global.GetConfig())
	if dbErr != nil {
		panic(dbErr)
	}
	
	mainLoopErr := MainLoop(db)
	if mainLoopErr != nil {
		panic(mainLoopErr)
	}
}