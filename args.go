package main

import (
	"github.com/jessevdk/go-flags"
	"strings"
)

func ParseArgs(args []string) (MavenPushCommand, error) {
	var command MavenPushCommand
	parser := flags.NewParser(&command, flags.None)
	_, err := parser.ParseArgs(args)
	if err != nil {
		return command, err
	}
	return command, nil
}

func RemoveMavenArgs(args []string) []string {
	for i, arg := range args {
		if strings.HasPrefix(arg, "--maven-") {
			return RemoveMavenArgs(removeArg(args, arg, i))
		}
	}
	return args
}

func removeArg(args []string, arg string, i int) []string {
	argsRemoved := 2
	if strings.Contains(arg, "=") {
		argsRemoved = 1
	}
	return append(args[0:i], args[i+argsRemoved:]...)
}
