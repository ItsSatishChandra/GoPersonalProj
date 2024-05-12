package main

import (
	"log"
	"website/backend"
)

func main() {
	err := backend.Server()
	if err != nil {
		log.Fatal("error running backend: ", err)
	}
}
