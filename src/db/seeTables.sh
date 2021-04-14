#!/bin/bash
echo "type 'USE ftvBackEnd'"
winpty docker exec -it mysqldb mysql -uroot -p
