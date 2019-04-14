#!/usr/bin/env bash

NAME=${1:-echoMock}
REGION=${2:-eu-west-3}

echo "Building ... "
GOOS=linux GOARCH=amd64 go build -o hello main.go
#hello is the value of the Lambda handler
zip hello.zip hello

echo "Uploading the new version"
aws lambda update-function-code --function-name ${NAME} --region ${REGION} --publish --zip-file fileb://hello.zip

rm hello
rm hello.zip

echo "Set the Lambda->Concurrency to a reserved 1 then run  a benchmark to simulate the users:"
echo "watch -n0 ab -n 2000 -c 6 -m GET -I http://<LB>.elb.amazonaws.com/"
