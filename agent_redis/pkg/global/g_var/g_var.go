package g_var

import "sync"

type globalVar struct {
	ID string
}

var once sync.Once
var g globalVar

func InitGlobalVar(id string) {
	once.Do(func(){
		g.ID = id
	})
}

func GetID() string {
	return g.ID
}


