package main

import (
	"database/sql"
	"time"
	"log"

	"sched/pkg/global"
	"sched/pkg/odbc/dsn"
)

func runSched(list []NextSched) (succList []int, failList []struct { id int; err error}) {
	if len(list) == 0 {
		return
	}

	failList = make([]struct{id int; err error}, 0)
	succList = make([]int, 0)

	conf := global.GetConfig()

	odbc_dsn := dsn.GenDsn(conf.OdbcDriver,
		conf.DBConfig.Host, 
		conf.DBConfig.User, 
		conf.DBConfig.Password, 
		conf.DBConfig.DbName, 
		conf.DBConfig.Port)

	exection := global.Execution()

	for _,sched := range list {
		err := exection.Run(sched.SchedName, odbc_dsn)

		if err != nil {
			log.Println(err.Error())
			failList = append(failList, struct{id int; err error}{id : sched.Id, err : err})
		}else {
			succList = append(succList, sched.Id)
		}
	}

	return
}


func MainLoop(db *sql.DB) error {
	for {
		time.Sleep(time.Millisecond * 50)

		schedList,err := DbSelectNextSched(db)		
		if(err != nil) {
			log.Println(err.Error())
			continue
		}

		succ, fail := runSched(schedList)
		
		for _, succId := range succ {
			DbInsertSchedLog(db, succId, "succ", "")
		}

		for _,failed := range fail {
			DbInsertSchedLog(db, failed.id, "fail", failed.err.Error())
		}
		
		succ = nil
		fail = nil		
	}
}