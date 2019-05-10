# flavor2go Flavor API

Loosely based on [go-vue-starter](https://github.com/markcheno/go-vue-starter/)

## running with docker

Install dep to manage the go dependencies and fetch them:

```sh
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

Then run docker-compose like:

```sh
yes | docker-compose rm && docker-compose up --build
```

If you want to run the locally build go app do:

```sh
go build && ./flavor2go -l 2999
```

This will listen on port 2999.

## development

Install dependencies:

```sh
go get -u ./...
```

## db migration

Go get the migrate CLI.

```sh
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
```

Run the migrations:

```sh
migrate -source file://migrations -database postgres://gusta:changeme@localhost:5432/gusta?sslmode=disable up
```

Create a new migration (e. g. for recipes). This will create up & down files for recipes in the migrations directory.

```sh
cd migrations
migrate create  -ext sql -dir . -seq -digits 4 recipes
```

To drop everything and start new do:

```sh
migrate -source file://migrations -database postgres://gusta:changeme@localhost:5432/gusta?sslmode=disable drop
```

Note: This will probably not be it. Lots of strange behaviors. E. g. using the above command with `-dir migrations/` fails in reading the existing files, because it can not convert the path to int. Too bad.
