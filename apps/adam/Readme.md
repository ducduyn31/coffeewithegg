# Adam Project

## Introduction

This project is created as a gateway and coordinator for coffeewithegg (CWE) project. It provides the following features:

- GraphQL endpoints for `Project` - `Project` refers to my personal projects, such as this one.
- GraphQL endpoints for `Infrastucture` - `Infrastructure` refers to the deployment plan that I use to host my projects, such as this one.

## Pre-requisites

- [Docker](https://www.docker.com/)
- [NX](https://nx.dev/)
- [Yarn](https://yarnpkg.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Golang](https://golang.org/)
- [Redis](https://redis.io/) - Only required for `Infrastructure` service

## Getting Started

This project is located under coffeewithegg monorepos, there are 2 ways to start this project:

### 1. Local dev machine

### Install

Under `coffeewithegg/` folder (root folder), run the following command:

```bash
yarn install
```

This will install the necessary libraries for the project, including [NX](https://nx.dev/), which will be used to run the project in dev environment.

### Run

Adam uses [Postgres](https://www.postgresql.org/) as its database. To start the database automatically, run the following command:

```bash
docker-compose up -d db
```

Or you can start the database manually on port `5432` with the following credentials:

```
username: postgres
password: password
```

Adam uses [Redis](https://redis.io/) for SlokyAPI - `Infrastructure` service. To start the cache automatically, run the following command:

```bash
docker-compose up -d redis
```

Adam uses [Viper](https://github.com/spf13/viper) as the configuration management. This configuration can be changed in `apps/adam/config.yml` file or can be overridden by environment variables.

To start the project, run the following command:

```bash
nx exec adam:serve
```

This will automatically run migrations in a similar way to [Liquibase](https://www.liquibase.org/). The migration files are located under `apps/adam/migrations` folder.

### Test

```bash
nx exec adam:test
```

### 2. Docker

### Build

```bash

```
