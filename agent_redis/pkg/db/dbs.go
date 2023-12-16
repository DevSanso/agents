package db

import (
	"context"
	"fmt"
)

func getDbSizeDb(rcc *redisClientCommander, ctx context.Context, dbname int) (int64, error) {
	script := fmt.Sprintf("redis.call('select',%d); return redis.call('dbsize')", dbname)
	cmd := rcc.client.Do(ctx, "eval", script, "0")
	if cmd.Err() != nil {
		return -1, cmd.Err()
	}	
	return cmd.Int64()
}
func getDbSizeDbOld(rcc *redisClientCommander, ctx context.Context, dbname int) (int64, error) {
	conn := rcc.client.Conn()
	defer conn.Close()

	selectCmd := conn.Pipeline().Do(ctx, "select ", dbname)
	if selectCmd.Err() != nil {
		return -1, selectCmd.Err()
	}

	sizeCmd := conn.DBSize(ctx)
	if sizeCmd.Err() != nil{
		return -1, sizeCmd.Err()
	}

	return sizeCmd.Val(), nil
}
func (rcc *redisClientCommander) GetDbSize(ctx context.Context, dbname int) (int64, error) {
	if rcc.dbVersion >= 2.6 {
		return getDbSizeDb(rcc, ctx, dbname)
	} else {
		return getDbSizeDbOld(rcc, ctx, dbname)
	}
}