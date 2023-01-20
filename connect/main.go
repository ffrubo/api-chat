package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
        fmt.Printf("Processing Lambda request %s. \n", request.RequestContext.RequestID)
        fmt.Printf("Connection ID: %s. \n", request.RequestContext.ConnectionID)
        fmt.Printf("Body size = %d. \n", len(request.Body))

        fmt.Println("Headers:")
        for k, v := range request.Headers {
                fmt.Printf("  %s: %s", k, v)
        }

        headers := map[string]string{
                "Content-Type": "application/json",
        }

        return events.APIGatewayProxyResponse{
                StatusCode: 200,
                Headers: headers,
                Body: request.Body,
                IsBase64Encoded: false,
        }, nil
}

func main() {
        lambda.Start(HandleRequest)
}
