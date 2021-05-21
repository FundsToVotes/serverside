#!/bin/bash
source ../gateway/export_env_variables.sh

GOOS=linux go build .
scp -i "~/.ssh/new_funds_to_votes_aws.pem" api_call $REMOTE_SERVER_LOGIN:~