package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	foo := os.Getenv("SUPABASE_URL")
	bar := os.Getenv("SUPABASE_KEY")
	message := fmt.Sprintf("Hello World %s! fod: %s bar: %s", event.Name, foo, bar)
	return &message, nil
}

func main() {
	// supabase.Connect()
	lambda.Start(HandleRequest)
}
