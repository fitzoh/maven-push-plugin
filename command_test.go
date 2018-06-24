package main

import (
	"testing"
)

var (
	baseConfig = MavenConfig{
		RepoUrl:      "config-url",
		GroupId:      "config-group",
		ArtifactId:   "config-id",
		Version:      "config-version",
		Classifier:   "config-classifier",
		Extension:    "config-extension",
		RepoUsername: "config-user",
		RepoPassword: "config-password",
	}
	overriddenConfig = MavenConfig{
		RepoUrl:      "comand-url",
		GroupId:      "comand-group",
		ArtifactId:   "comand-id",
		Version:      "comand-version",
		Classifier:   "comand-classifier",
		Extension:    "comand-extension",
		RepoUsername: "comand-user",
		RepoPassword: "comand-password",
	}

	emptyCommand = MavenPushCommand{
		RepoUrl:      "",
		GroupId:      "",
		ArtifactId:   "",
		Version:      "",
		Classifier:   "",
		Extension:    "",
		RepoUsername: "",
		RepoPassword: "",
	}

	overrideCommand = MavenPushCommand{
		RepoUrl:      "comand-url",
		GroupId:      "comand-group",
		ArtifactId:   "comand-id",
		Version:      "comand-version",
		Classifier:   "comand-classifier",
		Extension:    "comand-extension",
		RepoUsername: "comand-user",
		RepoPassword: "comand-password",
	}
)

func TestMergeKeepOriginal(t *testing.T) {
	config := emptyCommand.Merge(baseConfig)
	if config != baseConfig {
		t.Fail()
	}
}

func TestMergeOverride(t *testing.T) {
	config := overrideCommand.Merge(baseConfig)
	if config != overriddenConfig {
		t.Fail()
	}
}
