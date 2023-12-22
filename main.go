package main

import (
	"os"

	"github.com/swlee3306/oskindbeat/cmd"

	_ "github.com/swlee3306/oskindbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
