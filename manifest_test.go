package main

import (
	"testing"
)

func TestParseSimpleManifest(t *testing.T) {
	config, err := ExtractMavenConfigFromManifest("testdata/simple-manifest.yml")
	if err != nil {
		t.Fatal("failed to parse manifest", err)
	}
	if config != simpleConfig {
		t.Errorf("\nmanifest config doesn't match\nexpected %+v\ngot %+v", simpleConfig, config)
	}
}

func TestParseAllMavenKeysManifest(t *testing.T) {
	config, err := ExtractMavenConfigFromManifest("testdata/all-maven-keys-manifest.yml")
	if err != nil {
		t.Fatal("failed to parse manifest", err)
	}
	if config != allMavenKeysConfig {
		t.Errorf("\nmanifest config doesn't match\nexpected %+v\ngot %+v", simpleConfig, config)
	}
}
