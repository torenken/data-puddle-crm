package event

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Response(data any, statusCode int) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(data)
	headers := map[string]string{"Content-Type": "application/json"}

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}

func FailedResponse(err error) (events.APIGatewayProxyResponse, error) {
	return Response(err, 500)
}
