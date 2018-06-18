package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Manifest struct {
	Applications []Application `yaml:"applications"`
}

type Application struct {
	Name        string      `yaml:"name"`
	MavenConfig MavenConfig `yaml:"maven"`
}

func ParseManifest(f string) (Manifest, error) {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return Manifest{}, fmt.Errorf("failed to read manifest file %s, %+v", f, err)
	}
	var manifest Manifest
	err = yaml.Unmarshal(raw, &manifest)
	if err != nil {
		return Manifest{}, fmt.Errorf("failed to umarshall manifest file %s, %+v", f, err)
	}
	if numApplications := len(manifest.Applications); numApplications != 1 {
		return Manifest{}, fmt.Errorf("single application manifest required, %d found", numApplications)
	}
	manifest.Applications[0].MavenConfig.SetDefaults()
	return manifest, nil
}
