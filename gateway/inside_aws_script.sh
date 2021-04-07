#!/bin/bash
docker stop containerapi
docker rm -f containerapi
docker pull jhoupps/deployapi
docker run -d -p 443:443 \
--name containerapi \
--env TLSCERT="/etc/letsencrypt/live/api.jayhouppermans.me/fullchain.pem" \
--env TLSKEY="/etc/letsencrypt/live/api.jayhouppermans.me/privkey.pem" \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
jhoupps/deployapi