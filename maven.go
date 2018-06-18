package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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

func (config *MavenConfig) SetDefaults() {
	if len(config.Extension) == 0 {
		config.Extension = "jar"
	}
}

func DownloadArtifact(url string, destination string, username string, password string) error {
	output, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create temp dir, %+v", err)
	}
	defer output.Close()

	fmt.Printf("downloading artifact from %s\n", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request, %+v", err)
	}
	if len(username) != 0 {
		request.SetBasicAuth(username, password)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to download artifact from maven repository, %+v", err)
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("received status code %d from maven repository: %s\n", response.StatusCode, response.Status)
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("failed to save artifact to temp file, %+v", err)
	}
	return nil
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
