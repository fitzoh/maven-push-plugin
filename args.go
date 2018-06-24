package main

import (
	"github.com/jessevdk/go-flags"
)

func ParseArgs(args []string) (MavenPushCommand, error) {
	var command MavenPushCommand
	parser := flags.NewParser(&command, flags.Default|flags.IgnoreUnknown)
	_, err := parser.ParseArgs(args)
	if err != nil {
		println(err)
		return command, err
	}
	return command, nil
}
