#!/bin/bash
docker run -d \
-p 3306:3306 \
--name mysqldb \
-e MYSQL_ROOT_PASSWORD=ftvInternal123 \
-e MYSQL_DATABASE=ftvBackEnd \
mysql
