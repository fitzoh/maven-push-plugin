package main

import (
	"code.cloudfoundry.org/cli/command/v2"
)

type MavenPushCommand struct {
	RepoUrl      string `long:"maven-repo-url"`
	GroupId      string `long:"maven-group-id"`
	ArtifactId   string `long:"maven-artifact-id"`
	Version      string `long:"maven-version"`
	Classifier   string `long:"maven-classifier"`
	Extension    string `long:"maven-extension"`
	RepoUsername string `long:"maven-repo-username"`
	RepoPassword string `long:"maven-repo-password"`
	Push         v2.V2PushCommand
}

func (command MavenPushCommand) ManifestPath() string {
	if command.Push.PathToManifest == "" {
		return "manifest.yml"
	}
	return string(command.Push.PathToManifest)
}

func (command MavenPushCommand) Merge(config MavenConfig) MavenConfig {
	if command.RepoUrl != "" {
		config.RepoUrl = command.RepoUrl
	}
	if command.GroupId != "" {
		config.GroupId = command.GroupId
	}
	if command.ArtifactId != "" {
		config.ArtifactId = command.ArtifactId
	}
	if command.Version != "" {
		config.Version = command.Version
	}
	if command.Classifier != "" {
		config.Classifier = command.Classifier
	}
	if command.Extension != "" {
		config.Extension = command.Extension
	}
	if command.RepoUsername != "" {
		config.RepoUsername = command.RepoUsername
	}
	if command.RepoPassword != "" {
		config.RepoPassword = command.RepoPassword
	}
	return config
}
