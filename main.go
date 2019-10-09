package main

import (
	"os"

	"github.com/millenovecento/mqttbeat/cmd"

	_ "github.com/millenovecento/mqttbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
