# Go Mongo

**go mongo modules** is a blazing fast go-mongodb boilerplate for inspired from [go module](https://github.com/amrebada/go-modules)

## What is included

1. `dotenv` support
2. [mgm](https://github.com/Kamva/mgm) database support with models
3. `mongodb` database
4. [gofiber](https://gofiber.io/) support for web server, middlewares ...
5. [Artillery](https://www.artillery.io/) load testing boilerplate
6. [Validation](https://github.com/go-ozzo/ozzo-validation) of DTO, requests.
7. [Static Analysis](https://staticcheck.dev/docs/getting-started) for clean code.

## Future Contribution

- Adding Benchmarks using Artillery load testing.
- 🔎 `Github Action` _comming soon_
- Middleware for Authentication and Autorization.
- Shift from `goAir` for Hot-Module Reload instead of `nodemon`.

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
- run lint: `make lint`

## Contribution

You can request new features by creating an [issue](https://github.com/Yasir900aslam/go_mongo_modules/issues), or submit a [pull request](https://github.com/Yasir900aslam/go_mongo_modules/pulls) with your contribution.
