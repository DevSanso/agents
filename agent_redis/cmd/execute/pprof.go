package main

import (
	"net/http"
	"net/http/pprof"
)

func makePprofServer(arg string) *http.Server {
	if arg != "-pprof" {
		return nil
	}
	
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/debug/pprof/", pprof.Index)
	serverMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	serverMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	serverMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	serverMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	
	server := &http.Server{}
	server.Handler = serverMux
	server.Addr = "127.0.0.1:22861"

	return server
}
