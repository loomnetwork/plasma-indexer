Plasma Indexer
===

Sample code of EVM contract events indexer and API. 

## Setup sample code
1. create a MYSQL database for storing events, default DB name is `plasma-indexer`
2. clone the probject 
```sh
git clone git@github.com:loomnetwork/plasma-indexer.git
```
3. deploy `LoomStore` contract to your cluster
4. call `set` method to make the contract emit an event
5. compile LoomStore indexer and Plasma API using the following command
```sh
make
```

## LoomStore indexer
LoomStore indexer queries events emitted from LoomStore contract and store them in MySQL database

## Run LoomStore indexer
```sh
./loomstore-indexer --db-password rootpassword --db-user root --block-height 5714082 \
 --read-uri http://your-non-validating-node.com/query --contract-address 0x8AE87cb755837c22Ec3E105144d88E9CE6769A62
```

## Plasma API
Plasma API provides API endpoint for querying specific events from MySQL database

## Run Plasma API
```sh
vi plasma.yaml // edit config file
./plasma-api
```
then go to http://localhost:3333/loomstore_events to check the fetched events

## Working example 

This repository contains solidity contract `LoomStore` along with its abi.
To demonstrate fetching contract events, the following steps have been done
1. `LoomStore` has been deployed on extdev cluster with address `0x8AE87cb755837c22Ec3E105144d88E9CE6769A62`
2. `LoomStore` has a method `set` that, once it's called, emits `NewValueSet` event
3. `set` is called at block height 6714083 and the `NewValueSet` event has been emitted, please check
http://extdev-plasma-us1.dappchains.com/query/contractevents?fromBlock=6714083&toBlock=6714083
4. `LoomStore indexer` is used to scan extdev cluster and store `NewValueSet` events in MySQL database
5. `Plasma API` provides an API endpoint for easily querying contract events 