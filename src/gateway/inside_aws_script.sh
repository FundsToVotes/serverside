#!/bin/bash
source export_env_variables.sh

docker stop ftvgateway
docker rm -f ftvgateway
docker pull jhoupps/ftvgateway
docker run -d -p 443:443 -p 49152:3306 \
--name ftvgateway \
--env TLSCERT="/etc/letsencrypt/live/api.fundstovotes.info/fullchain.pem" \
--env TLSKEY="/etc/letsencrypt/live/api.fundstovotes.info/privkey.pem" \
--env GATEWAY_PORT=$GATEWAY_PORT \
--env SERVERSIDE_APP_PROPUBLICA_CONGRESS_API_KEY=$SERVERSIDE_APP_PROPUBLICA_CONGRESS_API_KEY \
--env SERVERSIDE_OPENSECRETS_API_KEY=$SERVERSIDE_OPENSECRETS_API_KEY \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
jhoupps/ftvgateway


