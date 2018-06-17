package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"fmt"
)

type MavenPushPlugin struct {
}

type Manifest struct {
	Applications []Application `yaml:"applications"`
}

type Application struct {
	Name        string      `yaml:"name"`
	MavenConfig MavenConfig `yaml:"maven"`
}

func ParseManifest(f string) Manifest {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Printf("failed to read manifest file %s, %+v\n", f, err)
		os.Exit(1)
	}
	var manifest Manifest
	err = yaml.Unmarshal(raw, &manifest)
	if err != nil {
		fmt.Printf("failed to umarshall manifest file %s, %+v\n", f, err)
		os.Exit(1)
	}
	setManifestDefaults(manifest)
	return manifest
}

func setManifestDefaults(manifest Manifest) {
	for i := 0; i < len(manifest.Applications); i++ {
		config := &manifest.Applications[i].MavenConfig
		if len(config.Extension) == 0 {
			config.Extension = "jar"
		}
	}
}

func (c *MavenPushPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if len(args) == 0 || args[0] != "maven-push" {
		os.Exit(0)
	}
	flags := flag.NewFlagSet("maven-push", flag.ExitOnError)
	manifestPath := flags.String("f", "manifest.yml", "Path to manifest")
	flags.Parse(args[1:])

	fmt.Printf("using manifest file %s\n", *manifestPath)
	manifest := ParseManifest(*manifestPath)

	if numApplications := len(manifest.Applications); numApplications != 1 {
		fmt.Printf("single application manifest required, %d found", numApplications)
		os.Exit(1)
	}
	config := manifest.Applications[0].MavenConfig

	artifactDir, err := ioutil.TempDir("", "cf-maven-push")
	if err != nil {
		fmt.Printf("failed to create temp dir, %+v", err)
		os.Exit(1)
	}
	defer os.Remove(artifactDir)
	artifactFile := artifactDir + "/artifact"

	DownloadArtifact(BuildArtifactUrl(config), artifactFile, config.RepoUsername, config.RepoPassword)

	args = append(args, "-p", artifactFile)
	args[0] = "push"

	cliConnection.CliCommand(args...)

}

func (c *MavenPushPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "maven-push",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "maven-push",
				HelpText: "Download and push an application based on maven coordinates defined in the manifest",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "cf maven-push [-f MANIFEST_PATH]",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(MavenPushPlugin))
}
