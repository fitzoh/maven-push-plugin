#!/usr/bin/env bash

go test
GOOS=linux GOARCH=amd64 go build -o artifacts/maven-push.linux64
GOOS=linux GOARCH=386 go build -o artifacts/maven-push.linux32
GOOS=windows GOARCH=amd64 go build -o artifacts/maven-push.win64
GOOS=windows GOARCH=386 go build -o artifacts/maven-push.win32
GOOS=darwin GOARCH=amd64 go build -o artifacts/maven-push.osx
