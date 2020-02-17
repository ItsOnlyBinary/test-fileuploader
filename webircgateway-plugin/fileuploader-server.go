package main

// symlink this file into $GOPATH/src/github.com/kiwiirc/webircgateway/plugins

import (
	"fmt"
	"sync"

	"github.com/itsonlybinary/test-fileuploader/server"
	"github.com/kiwiirc/webircgateway/pkg/webircgateway"
)

func Start(gateway *webircgateway.Gateway, pluginsQuit *sync.WaitGroup) {
	fmt.Println("fileuploader-server start")
	gateway.Log(1, "Starting fileuploader-server plugin. webircgateway version: %s", webircgateway.Version)

	go func() {
		defer pluginsQuit.Done()
		server.RunServer(gateway.HttpRouter, "fileuploader.config.toml")
	}()
}
