# Go Mongo

**go mongo modules** is a go-mongodb boilerplate for inspired from [go module](https://github.com/amrebada/go-modules)

## What is included

1. `dotenv` support
2. [mgm](https://github.com/Kamva/mgm) database support with modelsn
3. `mongodb` database
4. [gofiber](https://gofiber.io/) support for web server, middlewares ...
5. [Artillery](https://www.artillery.io/) load testing boilerplate
6. 🔎 `swagger generating` _comming soon_
7. 🔎 `Github Action` _comming soon_

## Prequisites

- nodemon `npm install -g nodemon`
- [go](https://golang.org/doc/install)
- make

## Optional

- Artillery `npm install -g artillery`

## Scripts

- development: `make dev`
- rename name of the project: `make rename name=<new project name>`
- build for linux: `make build`
- build for mac: `make build_mac`
- run database migration: `make migrate`
- run tests: `make test`

## Contribution

You can request new features by creating an [issue](https://github.com/Yasir900aslam/go_mongo_modules/issues), or submit a [pull request](https://github.com/Yasir900aslam/go_mongo_modules/pulls) with your contribution.
