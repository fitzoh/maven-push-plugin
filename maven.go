package main

import (
	"os"
	"net/http"
	"io"
	"fmt"
	"net/url"
	"strings"
	"log"
)

type MavenConfig struct {
	RepoUrl      string `yaml:"repo-url"`
	GroupId      string `yaml:"group-id"`
	ArtifactId   string `yaml:"artifact-id"`
	Version      string `yaml:"version"`
	Extension    string `yaml:"extension"`
	Classifier   string `yaml:"classifier"`
	RepoUsername string `yaml:"repo-username"`
	RepoPassword string `yaml:"repo-password"`
}

func DownloadArtifact(url string, destination string, username string, password string) {
	output, err := os.Create(destination)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	fmt.Printf("downloading artifact from %s\n", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("failed to create request, %+v", err)
	}
	if len(username) != 0 {
		request.SetBasicAuth(username, password)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("failed to download artifact from maven repository, %+v", err)
		os.Exit(1)
	}
	if response.StatusCode != 200 {
		fmt.Printf("received status code %d from maven repository: %s\n", response.StatusCode, response.Status)
		os.Exit(1)
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildArtifactName(config MavenConfig) string {
	if len(config.Classifier) == 0 {
		return fmt.Sprintf("%s-%s.%s", config.ArtifactId, config.Version, config.Extension)
	}
	return fmt.Sprintf("%s-%s-%s.%s", config.ArtifactId, config.Version, config.Classifier, config.Extension)
}

func BuildArtifactUrl(config MavenConfig) string {
	artifactUrl := config.RepoUrl
	validateUrl(config.RepoUrl)
	for _, fragment := range strings.Split(config.GroupId, ".") {
		artifactUrl = addPath(artifactUrl, fragment)
	}
	artifactUrl = addPath(artifactUrl, config.ArtifactId)
	artifactUrl = addPath(artifactUrl, config.Version)
	artifactUrl = addPath(artifactUrl, BuildArtifactName(config))
	validateUrl(artifactUrl)
	return artifactUrl
}

func addPath(base string, fragment string) string {
	if ! strings.HasSuffix(base, "/") {
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
