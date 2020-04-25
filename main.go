package main

import (
	"os"

	"github.com/fabiojaniolima/countbeat/cmd"

	_ "github.com/fabiojaniolima/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
