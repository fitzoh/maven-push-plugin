package main

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    MavenPushParameters
		wantErr bool
	}{
		{name: "no args", args: []string{"maven-push"}, want: MavenPushParameters{ManifestPath: "manifest.yml"}, wantErr: false},
		{name: "proxy args only", args: []string{"maven-push", "-i", "2"}, want: MavenPushParameters{ManifestPath: "manifest.yml"}, wantErr: false},
		{name: "manifest and proxy args", args: []string{"maven-push", "-f", "other-manifest.yml", "-i", "2"}, want: MavenPushParameters{ManifestPath: "other-manifest.yml"}, wantErr: false},
		{name: "app name no args", args: []string{"maven-push", "my-app"}, want: MavenPushParameters{ManifestPath: "manifest.yml"}, wantErr: false},
		{name: "app name proxy args only", args: []string{"maven-push", "my-app", "-i", "2"}, want: MavenPushParameters{ManifestPath: "manifest.yml"}, wantErr: false},
		{name: "app name manifest and proxy args", args: []string{"maven-push", "my-app", "-f", "other-manifest.yml", "-i", "2"}, want: MavenPushParameters{ManifestPath: "other-manifest.yml"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
