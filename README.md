# flavor2go Flavor API

Loosely based on [go-vue-starter](https://github.com/markcheno/go-vue-starter/)

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

Create a new migration (e. g. for recipes). This will create up & down files for vendors in the migrations directory.

```sh
cd migrations
migrate.exe create  -ext sql -dir . -seq -digits 4 recipes
```

Note: This will probably not be it. Lots of strange behaviors. E. g. using the above command with `-dir migrations/` fails in reading the existing files, because it can not convert the path to int. Too bad.
