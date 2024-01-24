package config

import (
	"errors"
)

var (
	ErrNotMatchSendType = errors.New("ErrNotMatchSendType")
)

type SenderConfig struct {
	SendType string

	SendConfig map[string]interface{}
}

func (sc *SenderConfig)MmapConfig() (filename string, size int, err error){
	if sc.SendType == "MMAP" {
		filename = sc.SendConfig["filename"].(string)
		size = sc.SendConfig["size"].(int)
		return
	}
	err = ErrNotMatchSendType
	return 
}

func (sc *SenderConfig)TcpConfig() (ip string, port int, err error) {
	if sc.SendType == "TCP" {
		ip = sc.SendConfig["ip"].(string)
		port = sc.SendConfig["port"].(int)
		return
	}
	err = ErrNotMatchSendType
	return 
}