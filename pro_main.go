//go:build pro

// +build: pro
package main

import (
	"log"

	"github.com/mathnogueira/ioc/commercial"
)

func main() {
	server, err := commercial.GetServer()
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
