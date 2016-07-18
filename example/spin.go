package main

import (
	"fmt"
	"time"

	"github.com/ddo/go-spin"
)

func main() {
	spinner := spin.New("")

	// spin every 100ms
	ticker := time.NewTicker(time.Millisecond * 100)

	go func() {
		for now := range ticker.C {
			fmt.Printf(" \r %s %s", spinner.Spin(), now)
		}
	}()

	// stop after 5s
	<-time.After(time.Second * 10)

	// this stop doesnt close the ticker.C channel
	ticker.Stop()

	fmt.Println("\ndone")
}
