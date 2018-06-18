package main

import (
	"testing"
)

func TestParseSimpleManifest(t *testing.T) {
	manifest, err := ParseManifest("testdata/simple-manifest.yml")
	if err != nil {
		t.Fatal("failed to parse manifest", err)
	}
	config := manifest.Applications[0].MavenConfig
	if config != simpleConfig {
		t.Errorf("\nmanifest config doesn't match\nexpected %+v\ngot %+v", simpleConfig, config)
	}
}

func TestParseAllMavenKeysManifest(t *testing.T) {
	manifest, err := ParseManifest("testdata/all-maven-keys-manifest.yml")
	if err != nil {
		t.Fatal("failed to parse manifest", err)
	}
	config := manifest.Applications[0].MavenConfig
	if config != allMavenKeysConfig {
		t.Errorf("\nmanifest config doesn't match\nexpected %+v\ngot %+v", simpleConfig, config)
	}
}
