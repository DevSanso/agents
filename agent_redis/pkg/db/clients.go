package db

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type RedisClientInfo struct {
	ID        int
	Addr      string
	LocalAddr string
	FD        int
	Name      string
	Age       int
	Idle      int
	Flags     string
	DB        int
	Sub       int
	PSub      int
	SSub      int
	Multi     int
	QBuf      string
	QBufFree  int
	ArgvMem   int
	MultiMem  int
	RBS       int
	RBP       int
	OBL       int
	OLL       int
	OMem      int
	TotMem    int
	Events    string
	Cmd       string
	User      string
	Redir     int
	Resp      int
	LibName   string
	LibVer    string
}

func parseRedisClientInfo(input string) (RedisClientInfo, error) {
	var client RedisClientInfo

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
			client.ID, _ = strconv.Atoi(value)
		case "addr":
			client.Addr = value
		case "laddr":
			client.LocalAddr = value
		case "fd":
			client.FD, _ = strconv.Atoi(value)
		case "name":
			client.Name = value
		case "age":
			client.Age, _ = strconv.Atoi(value)
		case "idle":
			client.Idle, _ = strconv.Atoi(value)
		case "flags":
			client.Flags = value
		case "db":
			client.DB, _ = strconv.Atoi(value)
		case "sub":
			client.Sub, _ = strconv.Atoi(value)
		case "psub":
			client.PSub, _ = strconv.Atoi(value)
		case "ssub":
			client.SSub, _ = strconv.Atoi(value)
		case "multi":
			client.Multi, _ = strconv.Atoi(value)
		case "qbuf":
			client.QBuf = value
		case "qbuf-free":
			client.QBufFree, _ = strconv.Atoi(value)
		case "argv-mem":
			client.ArgvMem, _ = strconv.Atoi(value)
		case "multi-mem":
			client.MultiMem, _ = strconv.Atoi(value)
		case "rbs":
			client.RBS, _ = strconv.Atoi(value)
		case "rbp":
			client.RBP, _ = strconv.Atoi(value)
		case "obl":
			client.OBL, _ = strconv.Atoi(value)
		case "oll":
			client.OLL, _ = strconv.Atoi(value)
		case "omem":
			client.OMem, _ = strconv.Atoi(value)
		case "tot-mem":
			client.TotMem, _ = strconv.Atoi(value)
		case "events":
			client.Events = value
		case "cmd":
			client.Cmd = value
		case "user":
			client.User = value
		case "redir":
			client.Redir, _ = strconv.Atoi(value)
		case "lib-name":
			client.LibName = value
		case "lib-ver":
			client.LibVer = value
		}
	}
	return client, nil
}

func convertStringListToRedisClientInfoList(input []string) ([]RedisClientInfo, error) {
	var result []RedisClientInfo
	
	for _, record := range input {
		client, err := parseRedisClientInfo(record)
		if err != nil {
			return nil, err
		}
		result = append(result, client)
	}

	return result, nil
}

func (rcc *redisClientCommander) GetClientList(ctx context.Context) ([]RedisClientInfo, error) {
	cmd := rcc.client.Do(ctx, "client", "list")
	if cmd.Err() != nil {
		return nil,cmd.Err()
	}

	lines := cmd.String()
	lines = strings.Replace(lines,"client list:","", 1)
	split := strings.Split(lines, "\n")
	return convertStringListToRedisClientInfoList(split)
}
