Plasma Indexer
===

Sample code of EVM contract scanner and API

## Setup

1. create a database for storing events, default name is `plasma-indexer`
2. compile LoomStore scanner and Plasma indexer using the following command
```sh
make deps
make
```

## LoomStore scanner
LoomStore scanner queries events emitted from LoomStore contract and store them in MySQL database

## Run LoomStore scanner
```sh
./loomsotre-scanner --db-password rootpassword --db-username root --block-height 5714082 --read-uri http://plasma.dappchains.com/query
```

## Plasma indexer
Plasma indexer provides HTTP endpoint for querying specific events from MySQL database

## Run Plasma indexer API
```sh
vi indexer.yaml // edit config file
./plasma-indexer
```