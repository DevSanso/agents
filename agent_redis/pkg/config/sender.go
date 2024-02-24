package config

import (
	"errors"
	"path/filepath"
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
		filename,_ = filepath.Abs(filename)
		size = int(sc.SendConfig["size"].(int64))
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