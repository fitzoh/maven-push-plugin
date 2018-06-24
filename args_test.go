package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	//move to testdata dir so that manifest validation works
	os.Chdir("testdata")
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{name: "no args", args: []string{"maven-push"}, want: "manifest.yml", wantErr: false},
		{name: "proxy args only", args: []string{"maven-push", "-i", "2"}, want: "manifest.yml", wantErr: false},
		{name: "manifest and proxy args", args: []string{"maven-push", "--no-route", "-f", "simple-manifest.yml", "-i", "2"}, want: "simple-manifest.yml", wantErr: false},
		{name: "manifest and proxy args and boolean", args: []string{"maven-push", "--no-route", "-f", "simple-manifest.yml", "-i", "2"}, want: "simple-manifest.yml", wantErr: false},

		{name: "app name no args", args: []string{"maven-push", "my-app"}, want: "manifest.yml", wantErr: false},
		{name: "app name proxy args only", args: []string{"maven-push", "my-app", "-i", "2"}, want: "manifest.yml", wantErr: false},
		{name: "app name manifest and proxy args", args: []string{"maven-push", "my-app", "-f", "simple-manifest.yml", "-i", "2"}, want: "simple-manifest.yml", wantErr: false},
		{name: "app name manifest and proxy args and boolean", args: []string{"maven-push", "my-app", "--no-route", "-f", "simple-manifest.yml", "-i", "2"}, want: "simple-manifest.yml", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ManifestPath(), tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
	//back to root dir so we don't mess up other tests
	os.Chdir("..")
}

func TestRemoveMavenArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{name: "empty", args: []string{}, want: []string{}},

		{name: "space", args: []string{"--maven-thing", "thing"}, want: []string{}},
		{name: "equals", args: []string{"--maven-thing=thing"}, want: []string{}},

		{name: "space with neighbors", args: []string{"keep1", "--maven-thing", "thing", "keep2"}, want: []string{"keep1", "keep2"}},
		{name: "equals with neighbors", args: []string{"keep1", "--maven-thing=thing", "keep2"}, want: []string{"keep1", "keep2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveMavenArgs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveMavenArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
