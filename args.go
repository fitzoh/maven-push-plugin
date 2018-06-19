package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

type MavenPushParameters struct {
	ManifestPath string
}

func ParseArgs(args []string) (MavenPushParameters, error) {
	params := MavenPushParameters{
		ManifestPath: "manifest.yml",
	}

	flags := flag.NewFlagSet("maven-push", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)
	manifestPath := flags.String("f", "manifest.yml", "Path to manifest")

	flags.Parse(prepareArgs(args))
	if len(*manifestPath) != 0 {
		params.ManifestPath = *manifestPath
	}
	return params, nil
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
