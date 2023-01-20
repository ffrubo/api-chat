package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

type payload struct {
  Action       string `json:"action"`
  ConnectionID string `json:"connectionId"`
  Data         string `json:"data"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        // Create config and get credentials
        cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
        if err != nil {
                log.Fatalf("unable to load SDK config, %v", err)
        }

        credentials, err := cfg.Credentials.Retrieve(context.TODO())
        if err != nil {
                log.Fatalf("unable to retrieve credentials, %v", err)
        }

        // Set body and compute payload hash

        payload := &payload{}
        err = json.Unmarshal([]byte(request.Body), payload)
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

        bodyString := fmt.Sprintf(`{"data":"%s", "source": "%s"}`, mockMessage.String(), "self-signed-post")
        body := strings.NewReader(bodyString)

        hash := sha256.New()
        hash.Write([]byte(bodyString))

        plHash := hex.EncodeToString(hash.Sum(nil))

        url := fmt.Sprintf(
                "%s/%s",
                os.Getenv("CONNECTIONS_URL"),
                payload.ConnectionID,
        )
        req, err := http.NewRequest(http.MethodPost, url, body)

        signer := v4.NewSigner()
        err = signer.SignHTTP(
                ctx,
                credentials,
                req,
                plHash,
                "execute-api",
                cfg.Region,
                time.Now(),
        )
        if err != nil {
                log.Fatalf("unable to sign request, %v", err)
        }


        res, err := http.DefaultClient.Do(req)
        if err != nil {
                log.Fatalf("unable to send request, %v", err)
        }

        defer res.Body.Close()
        if res.StatusCode != 200 {
                resBody, err := io.ReadAll(res.Body)
                if err != nil {
                        log.Printf("unable to read response body, %v", err)
                }

                log.Printf("message, %s", string(resBody))
                log.Fatalf("unexpected status code, %v\n", res.StatusCode)
        }

        headers := map[string]string{
                "Content-Type": "application/json",
        }

        return events.APIGatewayProxyResponse{
                StatusCode: 201,
                Headers: headers,
                Body: `{"data": "message sent successfully"}`,
                IsBase64Encoded: false,
        }, nil

}

func main() {
        lambda.Start(HandleRequest)
}
