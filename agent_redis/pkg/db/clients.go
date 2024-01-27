package db

import (
	"agent_redis/pkg/protos"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseRedisClientInfo(input string) (*protos.RedisClientInfo, error) {
	var client = &protos.RedisClientInfo{}

	fields := strings.Fields(input)

	for _, field := range fields {
		if field == "" {continue}
		keyValue := strings.SplitN(field, "=", 2)
		if len(keyValue) != 2 {
			return client, fmt.Errorf("invalid key-value pair: %s", input)
		}

		key, value := keyValue[0], keyValue[1]
		
		switch key {
		case "id":
			client.ID, _ = strconv.ParseInt(value, 10, 64)
		case "addr":
			client.Addr = value
		case "laddr":
			client.LocalAddr = value
		case "fd":
			client.FD, _ = strconv.ParseInt(value, 10, 64)
		case "name":
			client.Name = value
		case "age":
			client.Age, _ = strconv.ParseInt(value, 10, 64)
		case "idle":
			client.Idle, _ = strconv.ParseInt(value, 10, 64)
		case "flags":
			client.Flags = value
		case "db":
			client.DB, _ = strconv.ParseInt(value, 10, 64)
		case "sub":
			client.Sub, _ = strconv.ParseInt(value, 10, 64)
		case "psub":
			client.PSub, _ = strconv.ParseInt(value, 10, 64)
		case "ssub":
			client.SSub, _ = strconv.ParseInt(value, 10, 64)
		case "multi":
			client.Multi, _ = strconv.ParseInt(value, 10, 64)
		case "qbuf":
			client.QBuf = value
		case "qbuf-free":
			client.QBufFree, _ = strconv.ParseInt(value, 10, 64)
		case "argv-mem":
			client.ArgvMem, _ = strconv.ParseInt(value, 10, 64)
		case "multi-mem":
			client.MultiMem, _ = strconv.ParseInt(value, 10, 64)
		case "rbs":
			client.RBS, _ = strconv.ParseInt(value, 10, 64)
		case "rbp":
			client.RBP, _ = strconv.ParseInt(value, 10, 64)
		case "obl":
			client.OBL, _ = strconv.ParseInt(value, 10, 64)
		case "oll":
			client.OLL, _ = strconv.ParseInt(value, 10, 64)
		case "omem":
			client.OMem, _ = strconv.ParseInt(value, 10, 64)
		case "tot-mem":
			client.TotMem, _ = strconv.ParseInt(value, 10, 64)
		case "events":
			client.Events = value
		case "cmd":
			client.Cmd = value
		case "user":
			client.User = value
		case "redir":
			client.Redir, _ = strconv.ParseInt(value, 10, 64)
		case "lib-name":
			client.LibName = value
		case "lib-ver":
			client.LibVer = value
		}
	}
	return client, nil
}

func convertStringListToRedisClientInfoList(input []string) (*protos.RedisClientList, error) {
	var list []*protos.RedisClientInfo
	
	for _, record := range input {
		client, err := parseRedisClientInfo(record)
		if err != nil {
			return nil, err
		}
		list = append(list, client)
	}

	ret := &protos.RedisClientList{}
	ret.Clients = list
	ret.UnixEpoch = uint64(time.Now().Unix())

	return ret, nil
}

func (rcc *redisClientCommander) GetClientList(ctx context.Context) (*protos.RedisClientList, error) {
	cmd := rcc.client.Do(ctx, "client", "list")
	if cmd.Err() != nil {
		return nil,cmd.Err()
	}

	lines := cmd.String()
	lines = strings.Replace(lines,"client list:","", 1)
	split := strings.Split(lines, "\n")
	return convertStringListToRedisClientInfoList(split)
}
