package main

import (
	"strings"
	"github.com/jessevdk/go-flags"
	"code.cloudfoundry.org/cli/command/v2"
)

func ParseManifestPath(args []string) (string, error) {
	pushCommand := v2.V2PushCommand{PathToManifest: "manifest.yml"}
	parser := flags.NewParser(&pushCommand, flags.IgnoreUnknown)
	_, err := parser.ParseArgs(args)
	if err != nil {
		return "", err
	}
	return string(pushCommand.PathToManifest), nil
}

func prepareArgs(args []string) []string {
	return stripAppName(stripCommand(args))
}

func stripCommand(args []string) []string {
	return args[1:]
}

func stripAppName(args []string) []string {
	if len(args) > 1 && !strings.HasPrefix(args[0], "-") {
		return args[1:]
	}
	return args

}
