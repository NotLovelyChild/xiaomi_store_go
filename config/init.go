package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	Config *ini.File
	err    error
)

func init() {
	Config, err = ini.Load("./config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}
