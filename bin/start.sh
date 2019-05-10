#! /bin/bash

# FIXME: read from config
migrate -source file://migrations -database postgres://gusta:changeme@postgres:5432/gusta?sslmode=disable up

./flavor2go -h postgres

# eof
