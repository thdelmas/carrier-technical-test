#!/bin/bash

# Run the script to install & test the code

set -x

killall main

rm -vf go.mod go.sum main
go mod init 'theophile.world/carrier_test_palindrome' ;
go get 'github.com/gin-gonic/gin' &&
go build main.go &&
./main &

sleep 5 && curl -X POST 'http://127.0.0.1:8000/api/is_palindrome' -d '{"test_string": "tattarrattat"}' &&
sleep 2 && curl -X POST 'http://127.0.0.1:8000/api/is_palindrome' -d '{"test_string": "Yes ?"}'
