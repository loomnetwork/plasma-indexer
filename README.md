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

## LoomStore indexer
LoomStore indexer queries events emitted from LoomStore contract and store them in MySQL database

## Run LoomStore indexer
```sh
./loomsotre-indexer --db-password rootpassword --db-username root --block-height 5714082 --read-uri http://plasma.dappchains.com/query
```

## Plasma API
Plasma API provides API endpoint for querying specific events from MySQL database

## Run Plasma API API
```sh
vi plasma.yaml // edit config file
./plasma-api
```