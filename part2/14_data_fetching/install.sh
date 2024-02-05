#!/bin/bash

# Run the script to install & test the code

killall main

rm -vf go.mod go.sum main
go mod init 'theophile.world/carrier_data_fetching' ;
go build main.go && printf "\n\n\n" &&
./main 
