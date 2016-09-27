package server

import (
	"fmt"
)

type CommandLineOptions struct {
	Port int `short:"p" long:"port" description:"Http port" default:"1339"`
}

func Init(options CommandLineOptions) {
	fmt.Println(GetCurrentDir())

	fmt.Print(options.Port)
}
