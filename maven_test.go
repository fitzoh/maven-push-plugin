package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDownloadArtifact(t *testing.T) {
	message := "message"
	file := "testdata/artifact"
	defer os.Remove(file)

	DownloadArtifact(base64Url(message), file, "", "")

	contents, _ := ioutil.ReadFile(file)
	if got := string(contents); got != message {
		t.Errorf("TestDownloadArtifact() = %v, want %v", got, message)
	}
}

func base64Url(message string) string {
	base64message := base64.StdEncoding.EncodeToString([]byte(message))
	return fmt.Sprintf("http://httpbin.org/base64/%s", base64message)
}

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
