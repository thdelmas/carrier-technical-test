#!/bin/bash

# Run the script to install & test the code

set -x

killall main

go mod init 'theophile.world/carrier_test_user_creation' ;
go get 'github.com/gin-gonic/gin' &&
go get 'github.com/lib/pq' &&
go build main.go &&
./main &

sleep 2 && curl -v -X POST 'http://127.0.0.1:8000/api/users' -d '{"email": "contact@theophile.world"}'
