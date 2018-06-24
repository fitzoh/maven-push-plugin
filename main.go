package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type MavenPushPlugin struct {
}

func (c *MavenPushPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if len(args) == 0 || args[0] != "maven-push" {
		os.Exit(0)
	}

	command, err := ParseArgs(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("using manifest file %s\n", command.ManifestPath())
	config, err := ExtractMavenConfigFromManifest(command.ManifestPath())
	config = command.Merge(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	artifactDir, err := ioutil.TempDir("", "cf-maven-push")
	if err != nil {
		fmt.Printf("failed to create temp dir, %+v", err)
		os.Exit(1)
	}
	defer os.Remove(artifactDir)
	artifactFile := artifactDir + "/artifact"

	err = DownloadArtifact(config.ArtifactUrl(), artifactFile, config.RepoUsername, config.RepoPassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args = append(args, "-p", artifactFile)
	args[0] = "push"
	args = RemoveMavenArgs(args)

	fmt.Println("running: cf", strings.Join(args, " "))
	_, err = cliConnection.CliCommand(args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *MavenPushPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "maven-push",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 2,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "maven-push",
				HelpText: "Download and push an application based on maven coordinates defined in the manifest",

				UsageDetails: plugin.Usage{
					Usage: "cf maven-push [-f MANIFEST_PATH] [--maven-repo-url REPO_URL] [--maven-group-id GROUP_ID]\n   [--maven-artifact-id ARTIFACT_ID] [--maven-version VERSION] [--maven-classifier CLASSIFIER]\n   [--maven-extension EXTENSION] [--maven-repo-username REPO_USERNAME] [--maven-repo-password REPO_PASSWORD] <cf push flags>\n\n   cf maven-push APP_NAME [-f MANIFEST_PATH] [--maven-repo-url REPO_URL] [--maven-group-id GROUP_ID]\n   [--maven-artifact-id ARTIFACT_ID] [--maven-version VERSION] [--maven-classifier CLASSIFIER]\n   [--maven-extension EXTENSION] [--maven-repo-username REPO_USERNAME] [--maven-repo-password REPO_PASSWORD] <cf push flags>",
					Options: map[string]string{
						"f":                    "Path to manifest (default manifest.yml)",
						"-maven-repo-url":      "Maven repository to pull the artifact from (e.g. https://repo.maven.apache.org/maven2/)",
						"-maven-group-id":      "Maven groupId",
						"-maven-artifact-id":   "Maven artifactId",
						"-maven-version":       "Maven version",
						"-maven-classifier":    "Maven classifier (if any)",
						"-maven-extension":     "Maven extension (e.g. jar, zip) (default jar)",
						"-maven-repo-username": "Basic auth username when accessing Maven Repository (optional, default none",
						"-maven-repo-password": "Basic auth password when accessing Maven Repository (optional, default none",
					},
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(MavenPushPlugin))
}
