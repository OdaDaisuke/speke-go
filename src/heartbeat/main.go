package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Code int64 `json:"code"`
}

func Handler() (Response, error) {
	return Response{
		Code: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
