#!/bin/bash
docker stop ftvgateway
docker rm -f ftvgateway
docker run -d -p 80:80 \
--name ftvgateway \
--env TLSCERT="/c/Users/Jay/Documents/_My_Documents/_Current_Quarter_Academics/Info_490/serverside/fullchain.pem" \
--env TLSKEY="/c/Users/Jay/Documents/_My_Documents/_Current_Quarter_Academics/Info_490/serverside/privkey.pem" \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
jhoupps/ftvgateway

#WARNING - THIS DOESNT WORK

#--env TLSCERT="/etc/letsencrypt/live/api.jayhouppermans.me/fullchain.pem" \
#--env TLSKEY="/etc/letsencrypt/live/api.jayhouppermans.me/privkey.pem" \