#!/bin/bash

cat .env | sed 's/^/export /' | bash
go run main.go