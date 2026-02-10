#!/bin/bash

./network.sh down
./network.sh up
#echo "Waiting for CouchDB to initialize..."
#sleep 15
./network.sh createChannel -ca
./network.sh deployCC -ccp ../chaincode/ -ccn traderchaincode1 -c tradechannel1
./network.sh deployCC -ccp ../chaincode/ -ccn traderchaincode2 -c tradechannel2

# Init ledger
cd utils
./init_ledger.sh