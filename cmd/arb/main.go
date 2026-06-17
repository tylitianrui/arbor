package main

import (
	"log"
)

func main() {
	if err := Execute(); err != nil {
		// handle  err
		log.Fatal(err)
	}
}
