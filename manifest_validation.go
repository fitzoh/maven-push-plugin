package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ValidKeys = map[string]bool{
		"repo-url":      true,
		"group-id":      true,
		"artifact-id":   true,
		"version":       true,
		"extension":     true,
		"classifier":    true,
		"repo-username": true,
		"repo-password": true,
	}
)

type ValidationManifest struct {
	Applications []ValidationApplication `yaml:"applications"`
}

type ValidationApplication struct {
	MavenConfig map[string]string `yaml:"maven"`
}

func ValidateManifest(f string) error {
	manifest, err := parseValidationManifest(f)
	if err != nil {
		return err
	}
	if numApplications := len(manifest.Applications); numApplications != 1 {
		return fmt.Errorf("single application manifest required, %d found", numApplications)
	}
	config := manifest.Applications[0].MavenConfig
	for k := range config {
		if !ValidKeys[k] {
			return fmt.Errorf("invalid key found in maven configuration, %s", k)
		}
	}

	return nil
}

func parseValidationManifest(f string) (ValidationManifest, error) {
	manifest := ValidationManifest{}
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return manifest, fmt.Errorf("failed to open manifest file %s, %+v", f, err)
	}
	err = yaml.Unmarshal(raw, &manifest)
	if err != nil {
		return manifest, fmt.Errorf("failed to parse manifest file %s, %+v", f, err)
	}
	return manifest, nil
}
