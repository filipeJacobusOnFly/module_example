package main

import (
	"fmt"
	"os"

	"module_example/cmd"
	logs "module_example/src/logger"
)

func main() {
	logs.SetupLogger()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
