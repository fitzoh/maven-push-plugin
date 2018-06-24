package main

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{name: "no args", args: []string{"maven-push"}, want: "manifest.yml", wantErr: false},
		{name: "proxy args only", args: []string{"maven-push", "-i", "2"}, want: "manifest.yml", wantErr: false},
		{name: "manifest and proxy args", args: []string{"maven-push", "-f", "testdata/simple-manifest.yml", "-i", "2"}, want: "testdata/simple-manifest.yml", wantErr: false},
		{name: "app name no args", args: []string{"maven-push", "my-app"}, want: "manifest.yml", wantErr: false},
		{name: "app name proxy args only", args: []string{"maven-push", "my-app", "-i", "2"}, want: "manifest.yml", wantErr: false},
		{name: "app name manifest and proxy args", args: []string{"maven-push", "my-app", "-f", "testdata/simple-manifest.yml", "-i", "2"}, want: "testdata/simple-manifest.yml", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseManifestPath(tt.args)
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
