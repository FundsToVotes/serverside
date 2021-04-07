#!/bin/bash

GOOS=linux go build
go build .
docker build -t jhoupps/ftvgateway .
go clean