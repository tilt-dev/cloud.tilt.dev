#!/bin/bash

cd $(dirname $(dirname $0))
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/cloud-tilt-dev ./
