package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

type MavenConfig struct {
	RepoUrl      string `yaml:"repo-url"`
	GroupId      string `yaml:"group-id"`
	ArtifactId   string `yaml:"artifact-id"`
	Version      string `yaml:"version"`
	Classifier   string `yaml:"classifier"`
	Extension    string `yaml:"extension"`
	RepoUsername string `yaml:"repo-username"`
	RepoPassword string `yaml:"repo-password"`
}

func (config MavenConfig) ArtifactName() string {
	if len(config.Classifier) == 0 {
		return fmt.Sprintf("%s-%s.%s", config.ArtifactId, config.Version, config.Extension)
	}
	return fmt.Sprintf("%s-%s-%s.%s", config.ArtifactId, config.Version, config.Classifier, config.Extension)
}

func (config MavenConfig) ArtifactUrl() string {
	artifactUrl := config.RepoUrl
	validateUrl(config.RepoUrl)
	for _, fragment := range strings.Split(config.GroupId, ".") {
		artifactUrl = addPath(artifactUrl, fragment)
	}
	artifactUrl = addPath(artifactUrl, config.ArtifactId)
	artifactUrl = addPath(artifactUrl, config.Version)
	artifactUrl = addPath(artifactUrl, config.ArtifactName())
	validateUrl(artifactUrl)
	return artifactUrl
}

func addPath(base string, fragment string) string {
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}
	return base + url.PathEscape(fragment)
}

func validateUrl(toValidate string) {
	_, err := url.Parse(toValidate)
	if err != nil {
		log.Fatal(err)
	}
}
