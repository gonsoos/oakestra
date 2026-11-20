package main

import (
	"autoscaler/api"
	"autoscaler/logger"
	"log"
	"os"
	"strconv"
)

func setup() {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		debug = false
	}
	logger.Init(debug)
}

func StartLoopServer() error {
	return nil
}

func main() {
	setup()

	errc := make(chan error, 2)
	go func() { errc <- StartLoopServer() }()
	go func() { errc <- api.StartApiServer() }()

	// Block until one of the servers stops (expected to run forever).
	// log.Fatal is intentional here: it is the single top-level exit point.
	log.Fatal(<-errc)
}
