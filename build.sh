#!/bin/bash
export AWS_PROFILE=terraform
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

zip api_sourcer.zip bootstrap

aws s3 cp ./api_sourcer.zip s3://sourcers/api_sourcer.zip

aws lambda update-function-code --function-name  api_sourcer --s3-bucket sourcers --s3-key api_sourcer.zip