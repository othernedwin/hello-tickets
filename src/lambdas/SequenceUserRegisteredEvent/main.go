package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	helloTicketsEvents "hello-tickets/src/events"
	"hello-tickets/src/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/google/uuid"
)

func SequenceUserRegisteredEvent(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var userRegisteredData helloTicketsEvents.UserRegisteredData

	err := utils.UnmarshalAPIGatewayRequestBody(&userRegisteredData, request)

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to unmarshal api gateway request body data: %v", err)
	}

	// Generate a unique UUID for the user
	userID := "user-" + uuid.New().String()

	event := helloTicketsEvents.NewUserRegisteredEvent(userID, userRegisteredData.FirstName, userRegisteredData.LastName, userRegisteredData.Email, userRegisteredData.PhoneNumber, nil)

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
		PartitionKey: aws.String(userID),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("error putting record: %v", err)
	}

	body := map[string]interface{}{
		"userID": userID,
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
	lambda.Start(SequenceUserRegisteredEvent)
}
