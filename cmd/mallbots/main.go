package main

import (
	"math/rand"
	"os"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())
	if err := NewApp().Run(); err != nil {
		os.Exit(1)
	}
}
