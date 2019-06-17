#!/usr/bin/env bash

export DBHOST=localhost
export DBPASSWD=noredduser
export DBUSER=noredduser
go run -tags 'bindatafs' main.go