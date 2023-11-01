#!/bin/bash

GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

zip api_sourcer.zip bootstrap

aws s3 cp ./api_sourcer.zip s3://sourcers/api_sourcer.zip --profile terraform