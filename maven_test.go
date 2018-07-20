package main

import (
	"testing"
)

func TestArtifactName(t *testing.T) {
	type args struct {
		config MavenConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{simpleConfig}, "my-artifact-1.0.0.jar"},
		{"classifier", args{classifierConfig}, "my-artifact-1.0.0-javadoc.jar"},
		{"zip", args{zipConfig}, "my-artifact-1.0.0.zip"},
		{"allMavenKeys", args{allMavenKeysConfig}, "my-artifact-1.0.0-source.zip"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.config.ArtifactName(); got != tt.want {
				t.Errorf("ArtifactName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArtifactUrl(t *testing.T) {
	type args struct {
		config MavenConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{simpleConfig}, "https://repo.maven.apache.org/maven2/com/group/my/my-artifact/1.0.0/my-artifact-1.0.0.jar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.config.ArtifactUrl(); got != tt.want {
				t.Errorf("ArtifactUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
