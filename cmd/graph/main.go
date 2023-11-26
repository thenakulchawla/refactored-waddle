package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/thenakulchawla/refactored-waddle/pkg/graph"
)

// GraphCmd struct for graph-related subcommands
type CLI struct {
	BFS BFSCmd `cmd:"" help:"Perform BFS on a graph"`
}

// BFSCmd struct for running BFS
type BFSCmd struct {
	// In case I want to add any params
}

func (b *BFSCmd) Run() error {
	fmt.Println("Running BFS...")
	graph.RunBFS()
	return nil
}

func main() {
	cli := CLI{}

	// Kong setup with CLI struct
	ctx := kong.Parse(&cli,
		kong.Name("graph"),
		kong.Description("Graphs are necessary"),
	)

	// Execute the command
	err := ctx.Run()
	if err != nil {
		log.Fatal(err)
	}
}
