GoStarter
==========

> This starter kit is designed to get you up and running with a project structure optimized for developing RESTful API services in Go. It is an opinionated Go starter kit built on top of Chi, using battle tested libraries proven to provide a good foundation for a project written in Golang.

### Prerequisites

The codebase requires these development tools:

* Go compiler and runtime: 1.15.2 or greater.
* Docker Engine: 19.0.0 or greater.

### Go Dependencies

The project uses Go modules which should be vendored:

```shell
env GO111MODULE=on GOPRIVATE="github.com" go mod vendor
```

### Configuration

You may configure GoStarter using either a configuration file named .env, environment variables, or a combination of both. Environment variables are prefixed with GoStarter, and will always have precedence over values provided via file.

#### Server
```properties
SERVER_ADDRESS: localhost
SERVER_PORT: 8080
```

`ADDRESS` - `string`

Hostname to listen on.

`PORT` - `number`

Port number to listen on. Defaults to `8080`.

#### Database

```properties
DB_DRIVER: postgres
DB_HOST: dbhost
DB_USER: dbuser
DB_PASSWORD: dbpassword
DB_NAME: dbname
```

**UUID Type:** This project uses the `uuid-ossp` module to generate ids:

```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

**Migrations Note** Migrations are not applied automatically, so you will need to run them after you've built GoStarter.
* If built locally: `./go-starter migrate`
* Using Docker: `docker run --rm go-starter ./go-starter migrate`

### Start in Development

The recommended workflow is to use Docker and the compose file to build and run the service and resources.

```shell
docker compose -f d8t/docker-compose.dev.yml up --build
docker exec -it go-starter-db psql -U gostarter -d gostarter
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
docker exec -it go-starter go run main.go migrate
```

__Hot Reloading:__ GoStarter uses Reflex in development for hot reloading.

### Start In Production

```shell
make image
```

now run the image with (make sure you've run the migrations first)
```shell
docker run -it go-starter
```

### Run Tests
Running the tests locally requires a valid database connection. Configure `config.test.yaml` with the appropriate values.

```shell
make test-prepare
make test
```