#!/bin/bash
./build.sh
docker push jhoupps/deployapi
ssh ec2-user@ec2-54-245-136-18.us-west-2.compute.amazonaws.com < ./inside_aws_script.sh



#This line below will change by the person deploying. Once we have multiple developers, we can have a discussion about if there's
# a standard location we want to keep the key at. If not, we can just have different versions of this line for each of us
#ssh -i "/c/Users/Jay/.ssh/LegaciesAWSKey.pem" ec2-user@ec2-34-217-214-100.us-west-2.compute.amazonaws.com < ./inside_aws_script.sh