#!/bin/bash
source export_env_variables.sh

./build.sh
docker push jhoupps/ftvgateway
ssh -i "~/.ssh/new_funds_to_votes_aws.pem" $REMOTE_SERVER_LOGIN < ./inside_aws_script.sh