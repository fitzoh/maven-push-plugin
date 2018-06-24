package main

type MavenPushCommand struct {
	PathToManifest string `short:"f" description:"Path to manifest" default:"manifest.yml" `
	RepoUrl        string `long:"maven-repo-url"`
	GroupId        string `long:"maven-group-id"`
	ArtifactId     string `long:"maven-artifact-id"`
	Version        string `long:"maven-version"`
	Classifier     string `long:"maven-classifier"`
	Extension      string `long:"maven-extension"`
	RepoUsername   string `long:"maven-repo-username"`
	RepoPassword   string `long:"maven-repo-password"`
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
