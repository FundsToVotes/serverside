#!/bin/bash
docker stop ftvgateway
docker rm -f ftvgateway
docker pull jhoupps/ftvgateway
docker run -d -p 443:443 \
--name ftvgateway \
--env TLSCERT="/etc/letsencrypt/live/api.fundstovotes.info/fullchain.pem" \
--env TLSKEY="/etc/letsencrypt/live/api.fundstovotes.info/privkey.pem" \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
jhoupps/ftvgateway