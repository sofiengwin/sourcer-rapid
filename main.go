package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sofiengwin/sourcer/supabase"
)

type MyEvent struct {
	Name string `json:"name"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%s", err)
		log.Fatal("Error loading .env file")
	}
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
	supabase.Connect()
	// football.GetFixture()
	// lambda.Start(HandleRequest)
}
