#!/bin/bash
source export_env_variables.sh

GOOS=linux go build
go build .
docker build -t jhoupps/ftvgateway .
go clean