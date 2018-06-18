package main

import (
	"testing"
	"strings"
)

var (
	simpleConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Extension:  "jar",
	}
	classifierConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Classifier: "javadoc",
		Extension:  "jar",
	}
	zipConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Extension:  "zip",
	}
	complexConfig = MavenConfig{
		RepoUrl:    "https://repo.maven.apache.org/maven2/",
		GroupId:    "com.group.my",
		ArtifactId: "my-artifact",
		Version:    "1.0.0",
		Classifier: "complex",
		Extension:  "zip",
	}
)

func TestParseSimpleManifest(t *testing.T) {
	manifest, err := ParseManifest("testdata/simple-manifest.yml")
	if err != nil{
		t.Fatal("failed to parse manifest", err)
	}
	config := manifest.Applications[0].MavenConfig
	if config != simpleConfig {
		t.Errorf("\nmanifest config doesn't match\nexpected %+v\ngot %+v", simpleConfig, config)
	}
}

func TestParseBadMultiAppManifest(t *testing.T) {
	_, err := ParseManifest("testdata/bad-multi-app-manifest.yml")
	if err == nil || !strings.HasPrefix(err.Error(), "single application manifest required"){
		t.Errorf("expected error about multiple applications in manifest")
	}
}

