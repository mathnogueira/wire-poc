//go:build free

// +build: free
package main

import (
	"log"

	"github.com/mathnogueira/ioc/startup"
)

func main() {
	server, err := startup.GetServer()
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
