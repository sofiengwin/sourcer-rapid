package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sofiengwin/sourcer/supabase"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	foo := os.Getenv("FOO")
	bar := os.Getenv("BAR")
	message := fmt.Sprintf("Hello World %s! fod: %s bar: %s", event.Name, foo, bar)
	return &message, nil
}

func main() {
	supabase.Connect()
	// lambda.Start(HandleRequest)
}
