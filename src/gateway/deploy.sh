#!/bin/bash
./build.sh
docker push jhoupps/ftvgateway
ssh -i "~/.ssh/new_funds_to_votes_aws.pem" ec2-user@ec2-54-189-210-105.us-west-2.compute.amazonaws.com < ./inside_aws_script.sh