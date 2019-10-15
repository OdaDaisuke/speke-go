package main

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/OdaDaisuke/speke_go/src/copyProtection/config"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
}

type Event struct {
	Name            string `json:"name"`
	IsBase64encoded bool   `json:"isBase64Encoded"`
	Body            string `json:"body"`
}

func Handler(ctx context.Context, event Event) (Response, error) {
	c := config.NewConfig()
	cpixXml := event.Body
	if event.IsBase64encoded {
		sDec, err := base64.StdEncoding.DecodeString(cpixXml)
		if err != nil {
			log.Fatalf("Error while decoding cpix xml\n")
		}
		cpixXml = string(sDec)
	}
	response, err := NewResponseGenerator(c).Run(cpixXml)
	if err != nil {
		log.Fatalf("generating response: %s", err)
	}

	return Response{
		Body:            response,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
