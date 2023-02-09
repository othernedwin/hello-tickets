package utils

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func UnmarshalAPIGatewayRequestBody(v any, request events.APIGatewayProxyRequest) error {

	var bytes []byte = []byte(request.Body)

	if request.IsBase64Encoded {
		decodedBytes, err := base64.StdEncoding.DecodeString(request.Body)

		if err != nil {
			return err
		}

		bytes = decodedBytes
	}

	err := json.Unmarshal(bytes, v)

	if err != nil {
		return err
	}

	return nil
}
