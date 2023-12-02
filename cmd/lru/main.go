package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/thenakulchawla/refactored-waddle/pkg/lru"
)

type CLI struct {
	LRU LRUCmd `cmd:"" help:"Run LRU Cache"`
}

type LRUCmd struct {
}

func (l *LRUCmd) Run() error {
	fmt.Println("running lru")

	lru.RunLRU()
	return nil

}

func main() {
	cli := CLI{}

	// Kong setup with CLI struct
	ctx := kong.Parse(&cli,
		kong.Name("lru"),
		kong.Description("LRU Cache"),
	)

	// Execute the command
	err := ctx.Run()
	if err != nil {
		log.Fatal(err)
	}
}
