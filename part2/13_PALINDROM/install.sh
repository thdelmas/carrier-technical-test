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
sleep 2 && curl -X POST 'http://127.0.0.1:8000/api/is_palindrome' -d '{"test_string": "Yes ?"}' &&

cat <<EOF

1. Get a string
2. Iterate on the string characters by characters to the half, for each character check if it's opposite in the string is the same
3. 1st char compared to -1 (last one), 2 char compared to the -2, etc...
4. except for middle char in case of string with odd number of chars
5. if at least one comparison fails it's not a palindrome return fgalse, returns yes otherwise
6. return the result

EOF

