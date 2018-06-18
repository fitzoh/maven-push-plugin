package main

import (
	"strings"
	"testing"
)

func TestSimpleManifest(t *testing.T) {
	err := ValidateManifest("testdata/simple-manifest.yml")
	if err != nil {
		t.Fatal("simple manifest should validate")
	}
}

func TestAllMavenKeysManifest(t *testing.T) {
	err := ValidateManifest("testdata/all-maven-keys-manifest.yml")
	if err != nil {
		t.Fatal("all maven keys manifest should validate")
	}
}

func TestMissingManifest(t *testing.T) {
	err := ValidateManifest("testdata/doesntexist")
	if err == nil || !strings.HasPrefix(err.Error(), "failed to open manifest") {
		t.Fatal("expected error opening file")
	}
}

func TestInvalidManifest(t *testing.T) {
	err := ValidateManifest("testdata/invalid-yaml-manifest.yml")
	if err == nil || !strings.HasPrefix(err.Error(), "failed to parse manifest") {
		t.Fatal("expected error parsing yaml")
	}
}

func TestEmptyManifest(t *testing.T) {
	err := ValidateManifest("testdata/empty-manifest.yml")
	if err == nil || !strings.HasPrefix(err.Error(), "single application manifest required") {
		t.Fatal("expected error for missing application")
	}
}

func TestMultipleApplicationManifest(t *testing.T) {
	err := ValidateManifest("testdata/multi-app-manifest.yml")
	if err == nil || !strings.HasPrefix(err.Error(), "single application manifest required") {
		t.Fatal("expected error for missing application")
	}
}

func TestInvalidKey(t *testing.T) {
	err := ValidateManifest("testdata/invalid-key-manifest.yml")
	if err == nil || !strings.HasPrefix(err.Error(), "invalid key found in maven configuration") {
		t.Fatal("expected error for invalid maven key")
	}
}
