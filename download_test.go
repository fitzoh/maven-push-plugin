package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	message := "message"
	file := "testdata/artifact"
	defer os.Remove(file)

	DownloadFile(base64Url(message), file, "", "")

	contents, _ := ioutil.ReadFile(file)
	if got := string(contents); got != message {
		t.Errorf("TestDownloadArtifact() = %v, want %v", got, message)
	}
}

func base64Url(message string) string {
	base64message := base64.StdEncoding.EncodeToString([]byte(message))
	return fmt.Sprintf("http://httpbin.org/base64/%s", base64message)
}
