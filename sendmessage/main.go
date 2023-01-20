package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type payload struct {
  Action       string `json:"action"`
  Data         string `json:"data"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
        payload := &payload{}
        err := json.Unmarshal([]byte(request.Body), payload)
        if err != nil {
                log.Fatalf("unable to unmarshal payload, %v", err)
        }

        // Create mock message
        var mockMessage bytes.Buffer
        for i := range payload.Data {
                var char string
                if i%2 == 0 {
                        char = strings.ToUpper(fmt.Sprintf("%c", (payload.Data[i])))
                } else {
                        char = strings.ToLower(fmt.Sprintf("%c", (payload.Data[i])))
                }
                mockMessage.WriteString(char)
        }

        resBodyString := fmt.Sprintf(`{"data":"%s", "connectionId": "%s"}`, mockMessage.String(), request.RequestContext.ConnectionID)

        headers := map[string]string{
                "Content-Type": "application/json",
        }

        return events.APIGatewayProxyResponse{
                StatusCode: 200,
                Headers: headers,
                Body: resBodyString,
                IsBase64Encoded: false,
        }, nil

}

func main() {
        lambda.Start(HandleRequest)
}
