package web

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Response(data any, statusCode int) (events.APIGatewayProxyResponse, error) {
	if statusCode == http.StatusNoContent {
		return events.APIGatewayProxyResponse{StatusCode: statusCode}, nil
	}

	body, _ := json.Marshal(data)
	headers := map[string]string{"Content-Type": "application/json"}

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}
