package main

import (
	"os"

	"github.com/haringattu/lsbeat/cmd"

	_ "github.com/haringattu/lsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
