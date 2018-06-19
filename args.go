package main

import (
	"flag"
	"io/ioutil"
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
	flags.Parse(args[1:])
	if len(*manifestPath) != 0 {
		params.ManifestPath = *manifestPath
	}
	return params, nil
}
