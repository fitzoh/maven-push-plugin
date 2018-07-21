package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, destination string, username string, password string) error {
	output, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s, %+v", destination, err)
	}
	defer output.Close()

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request, %+v", err)
	}
	if len(username) != 0 {
		request.SetBasicAuth(username, password)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to download file from %s, %+v", url, err)
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("received status code %d from url %s: %s\n", response.StatusCode, url, response.Status)
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("failed to save response to file, %+v", err)
	}
	return nil
}
