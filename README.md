# AWS Lambda example

A small lambda function that simulates a heavy I/O operation ~150ms duration.

It would use an Application Load Balancer that will trigger lambda functions in 2 AWS AZs. 

## Requirements
* bash, zip
* Go 1.11+ (for go modules)
* AWS account with admin rights
* AWS CLI configured 

## AWS setup
* Create a lambda functions on AWS: GO 1.x runtime, from scratch called "echoMock". Leave the Handler: hello

* Edit the lamda function, add the "Application Load Balancer" from the left trigger list and press "create new". Use the following parameters: 
    * Name: hello-mock
    * internet facing with HTTP 80 listener
    * choose 2 random AZ zones
    * reuse or create a new security group that has the TCP 80 open port for all sources
    * Routing: echoMock with Lambda Function
    * Register targets: "echoMock" Lambda function with $LATEST version
    * get the DNS name for the APP Load Balancer from the EC2/Load Balancer page, eg: "hello-mock-157433221.eu-west-3.elb.amazonaws.com"
    * wait until you will get a response from 
    ```bash
    curl http://hello-mock-157433221.eu-west-3.elb.amazonaws.com
    ```
    `Hello from Lambda!`
    
    See the lambda info: 
```bash
aws lambda get-function --function-name echoMock --region eu-west-3
```

### Build & Publish
After building the Go binary we can use the ` lambda update-function-code` command to upload it. See the script:

```bash
./publish.sh [function-name] [aws-region]
```