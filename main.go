package main

import (
	"flag"

	"github.com/itsonlybinary/test-fileuploader/server"
)

func main() {
	var configPath = flag.String("config", "fileuploader.config.toml", "path to config file")
	flag.Parse()
	runCtx := server.NewRunContext(nil, *configPath)
	runCtx.Run()
}
