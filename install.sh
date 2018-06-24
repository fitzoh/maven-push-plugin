#!/usr/bin/env bash

go test
go build
cf install-plugin maven-push -f