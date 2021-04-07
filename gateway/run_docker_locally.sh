#!/bin/bash
docker stop webserver
docker rm -f webserver
docker run -d -p 80:80 \
--name webserver \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
legaciesdev/webserver


#--env TLSCERT="/etc/letsencrypt/live/api.jayhouppermans.me/fullchain.pem" \
#--env TLSKEY="/etc/letsencrypt/live/api.jayhouppermans.me/privkey.pem" \