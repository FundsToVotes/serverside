#!/bin/bash
docker stop ftvgateway
docker rm -f ftvgateway
docker run -d -p 80:80 \
--name ftvgateway \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
jhoupps/ftvgateway


#--env TLSCERT="/etc/letsencrypt/live/api.jayhouppermans.me/fullchain.pem" \
#--env TLSKEY="/etc/letsencrypt/live/api.jayhouppermans.me/privkey.pem" \