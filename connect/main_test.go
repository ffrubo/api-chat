package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	var ctx context.Context
	event := events.APIGatewayWebsocketProxyRequest{
		Body: "test",
		RequestContext: events.APIGatewayWebsocketProxyRequestContext{
			ConnectionID: "abc-123",
		},
	}

	res, err := HandleRequest(ctx, event)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	if res.Body != "test" {
		t.Errorf("Expected body to be 'test', got %s", res.Body)
	}

	if res.IsBase64Encoded != false {
		t.Errorf("Expected IsBase64Encoded to be false, got %t", res.IsBase64Encoded)
	}
}

