package global

import (
	"sync"
	"log"
	"sched/pkg/conf"
)

var global struct {
	once sync.Once

	config conf.Configure
	log *log.Logger
}
