package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	helloTicketsEvents "hello-tickets/src/events"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func SequenceEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var eventData []byte = []byte(request.Body)

	if request.IsBase64Encoded {
		decodedBytes, err := base64.StdEncoding.DecodeString(request.Body)

		if err != nil {
			return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to decode request api gateway request body data: %v", err)
		}

		eventData = decodedBytes
	}

	eventType := request.QueryStringParameters["eventType"]

	key, event, err := helloTicketsEvents.NewEvent(eventType, eventData, nil)

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to initialize event: %v", err)
	}

	data, err := json.Marshal(event)

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to marshal event data: %v", err)
	}

	region := os.Getenv("AWS_REGION")

	sess, _ := session.NewSession(&aws.Config{
		Region: &region,
	})

	kinesisClient := kinesis.New(sess)

	stream := os.Getenv("APP_STREAM_NAME")

	// Write the event to the Kinesis stream
	_, err = kinesisClient.PutRecord(&kinesis.PutRecordInput{
		StreamName:   &stream,
		Data:         data,
		PartitionKey: &key,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("error putting record in kinesis: %v", err)
	}

	body := map[string]interface{}{
		"partitionKey": key,
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to marshal response body: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonBody),
	}, nil
}

func main() {
	lambda.Start(SequenceEvent)
}
