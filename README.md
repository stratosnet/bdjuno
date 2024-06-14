# Callisto
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/forbole/bdjuno/Tests)](https://github.com/forbole/bdjuno/actions?query=workflow%3ATests)
[![Go Report Card](https://goreportcard.com/badge/github.com/forbole/bdjuno)](https://goreportcard.com/report/github.com/forbole/bdjuno)
![Codecov branch](https://img.shields.io/codecov/c/github/forbole/bdjuno/cosmos/v0.40.x)

Callisto (formerly BDJuno) is the [Juno](https://github.com/forbole/juno) implementation
for [Big Dipper](https://github.com/forbole/big-dipper).

It extends the custom Juno behavior by adding different handlers and custom operations to make it easier for Big Dipper
showing the data inside the UI.

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).

## Usage
To know how to setup and run Callisto, please refer to
the [docs website](https://docs.bigdipper.live/cosmos-based/parser/overview/).

## Testing
If you want to test the code, you can do so by running

```shell
$ make test-unit
```

**Note**: Requires [Docker](https://docker.com).

This will:
1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.


## Local runner

```shell
docker-compose -f docker-compose.local.yml up -d
```

NOTE: You need to perform first time after launch

1. Create tables for bdjuno in stratos_db

    You need to run the SQL queries that you can find inside the database/schema folder.


2. Load mesos genesis
```shell
curl https://raw.githubusercontent.com/stratosnet/stratos-chain-testnet/main/mesos-1/genesis.json > local/genesis.json
```

3. Launch hasura container
```shell
docker compose up -d hasura
docker-compose -f docker-compose.local.yml exec hasura sh
```

and execute the following lines
```shell
apt-get update && apt-get install -y curl bash
curl -L https://github.com/hasura/graphql-engine/raw/stable/cli/get.sh | bash

hasura metadata apply --skip-update-check --project ./hasura
```