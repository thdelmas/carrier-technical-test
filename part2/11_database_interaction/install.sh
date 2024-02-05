#!/bin/bash

# Run the script to install & test the code

set -x

killall main

go mod init 'theophile.world/carrier_test_database' ;
go get 'github.com/hashicorp/go-memdb' &&
go get 'github.com/gin-gonic/gin' &&
go build main.go &&
./main &

sleep 2 && curl 'http://127.0.0.1:8000/api/products'
