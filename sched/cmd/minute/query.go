package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)
const (
	ConfSchema = "sched_min_conf"
)

var (
	SchedNextQuery = fmt.Sprintf(`
	select id, scehd_name
		from %s.minute_sched 
	where now() - last_execute_time > 60
	`,ConfSchema)

	UpdateLastExcuteTimeQuery = fmt.Sprintf(`
	update %s.minute_sched 
		set last_execute_time = now()
	where id = ?`, ConfSchema)

	InsertSchedExecuteLog = fmt.Sprintf(`
	insert %s.minute_sched_log(id, execute_time, state, text) values(?, now(),?,?)
	`, ConfSchema)

)
type NextSched struct {
	Id int
	SchedName string
}

func DbSelectNextSched(db *sql.DB) ([]NextSched, error) {
	conn, err := db.Conn(context.Background())
	if conn != nil {
		return nil,err
	}
	defer conn.Close()

	timeoutCtx,cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	rows,rowErr := conn.QueryContext(timeoutCtx, SchedNextQuery)
	if rowErr != nil {
		return nil, rowErr
	}
	defer rows.Close()
	

	ret := make([]NextSched, 0)
	{
		var fetchErr error = nil
		for rows.Next() {
			temp := NextSched{}
			fetchErr = rows.Scan(&temp.Id, &temp.SchedName)

			if fetchErr != nil {
				break
			}

			ret = append(ret, temp)
		}

		if fetchErr != nil {
			ret = nil
			return nil,fetchErr
		}
	}

	return ret, nil
}

func DbUpdateLastExecuteTime(db *sql.DB, id int) error {
	conn, err := db.Conn(context.Background())
	if conn != nil {
		return err
	}
	defer conn.Close()

	timeoutCtx,cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()
	
	_, updateErr := conn.ExecContext(timeoutCtx, UpdateLastExcuteTimeQuery, id)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func DbInsertSchedLog(db *sql.DB, id int, state string, text string) error {
	conn, err := db.Conn(context.Background())
	if conn != nil {
		return err
	}
	defer conn.Close()

	timeoutCtx,cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	_, insertErr := conn.ExecContext(timeoutCtx, InsertSchedExecuteLog, id, state, text)
	if insertErr != nil {
		return insertErr
	}
	return nil
}