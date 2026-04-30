// Package main is the entry point for engram, a local-first knowledge and memory
// management tool that stores context chunks for AI-assisted development workflows.
//
// Personal fork: customized for my own dev workflow and learning purposes.
package main

import (
	"fmt"
	"os"

	"github.com/Gentleman-Programming/engram/cmd"
)

const version = "0.1.0"

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
