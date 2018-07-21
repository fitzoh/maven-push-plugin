package main

import (
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v2"
	"fmt"
	"path/filepath"
)

type MavenPushCommand struct {
	RepoUrl           string `long:"maven-repo-url"`
	GroupId           string `long:"maven-group-id"`
	ArtifactId        string `long:"maven-artifact-id"`
	Version           string `long:"maven-version"`
	Classifier        string `long:"maven-classifier"`
	Extension         string `long:"maven-extension"`
	RepoUsername      string `long:"maven-repo-username"`
	RepoPassword      string `long:"maven-repo-password"`
	RemoteManifestUrl string `long:"remote-manifest-url"`
	Push              v2.V2PushCommand
}

func (cmd *MavenPushCommand) ManifestPath() string {
	if cmd.Push.PathToManifest == "" {
		return "manifest.yml"
	}
	return string(cmd.Push.PathToManifest)
}

func (cmd *MavenPushCommand) Merge(config MavenConfig) MavenConfig {
	if cmd.RepoUrl != "" {
		config.RepoUrl = cmd.RepoUrl
	}
	if cmd.GroupId != "" {
		config.GroupId = cmd.GroupId
	}
	if cmd.ArtifactId != "" {
		config.ArtifactId = cmd.ArtifactId
	}
	if cmd.Version != "" {
		config.Version = cmd.Version
	}
	if cmd.Classifier != "" {
		config.Classifier = cmd.Classifier
	}
	if cmd.Extension != "" {
		config.Extension = cmd.Extension
	}
	if cmd.RepoUsername != "" {
		config.RepoUsername = cmd.RepoUsername
	}
	if cmd.RepoPassword != "" {
		config.RepoPassword = cmd.RepoPassword
	}
	return config
}

func (cmd *MavenPushCommand) ConfigureRemoteManifestIfPresent(tempDir string) error {
	if cmd.RemoteManifestUrl == "" {
		return fmt.Errorf("remote manifest url is not configured")
	}
	manifestFile := filepath.Join(tempDir, "manifest.yml")
	err := DownloadFile(cmd.RemoteManifestUrl, manifestFile, "", "")
	if err != nil {
		return err
	}
	cmd.Push.PathToManifest = flag.PathWithExistenceCheck(manifestFile)
	return nil
}
