package main

import (
	"log"

	raspberry "github.com/angelniebla/go-libp2p-uart-bridge/src/dragino"
)

func main() {
	_, err := raspberry.NewBridge(true)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 1000; i++ {
		raspberry.WriteBridge([]byte(fmt.Sprintf("hello: %v", i)))
	}
}
