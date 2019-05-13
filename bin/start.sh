#! /bin/bash

# FIXME: read from config
migrate -source file://migrations -database postgres://gusta:changeme@postgres:5432/gusta?sslmode=disable up

./go-api -h postgres

# eof
