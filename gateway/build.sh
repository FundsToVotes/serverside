#!/bin/bash

GOOS=linux go build
go build .
docker build -t jhoupps/deployapi .
go clean