package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
	"time"
)

// https://aws.amazon.com/blogs/networking-and-content-delivery/lambda-functions-as-targets-for-application-load-balancers/
//https://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html#respond-to-load-balancer
// https://github.com/aws/aws-lambda-go/blob/master/events/README_ALBTargetGroupEvents.md

func handleRequest(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	//simulate a heavy I/O call
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))

	return events.ALBTargetGroupResponse{StatusCode: 200, StatusDescription: "200 OK", IsBase64Encoded: false,Body:"Echo",Headers:map[string]string{}}, nil
}

func main() {
	lambda.Start(handleRequest)
}