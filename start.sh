#!/bin/sh

# Use when running in local development

export ACCOUNTID=
export TOKEN=
export SERVICE_NUMBER=

go run main.go handler.go
