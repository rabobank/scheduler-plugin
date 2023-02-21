#!/bin/bash

OUTPUT_DIR=$PWD/dist
mkdir -p ${OUTPUT_DIR}

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/linux_amd64/scheduler-plugin -ldflags "-X github.com/rabobank/scheduler-plugin/conf.VERSION=${VERSION} -X github.com/rabobank/scheduler-plugin/conf.COMMIT=${COMMIT}" .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT_DIR}/darwin_amd64/scheduler-plugin -ldflags "-X github.com/rabobank/scheduler-plugin/conf.VERSION=${VERSION} -X github.com/rabobank/scheduler-plugin/conf.COMMIT=${COMMIT}" .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${OUTPUT_DIR}/darwin_arm64/scheduler-plugin -ldflags "-X github.com/rabobank/scheduler-plugin/conf.VERSION=${VERSION} -X github.com/rabobank/scheduler-plugin/conf.COMMIT=${COMMIT}" .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_DIR}/windows_amd64/scheduler-plugin -ldflags "-X github.com/rabobank/scheduler-plugin/conf.VERSION=${VERSION} -X github.com/rabobank/scheduler-plugin/conf.COMMIT=${COMMIT}" .
