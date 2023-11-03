package supabase

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

func main() {
	// client, err := supabase.NewClient("API_URL", "API_KEY", nil)
	client, err := supabase.NewClient("https://xtsujlxfyspxskmpsreg.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Inh0c3VqbHhmeXNweHNrbXBzcmVnIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTMyODUzNTQsImV4cCI6MjAwODg2MTM1NH0.-iN2Pgxhq9Q0bLj28bwKAT-PdulumbQ8lFBp66qNEQI", nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()

	fmt.Printf("%s, %d, %s", data, count, err)
}

func Connect() {
	// client, err := supabase.NewClient("API_URL", "API_KEY", nil)
	client, err := supabase.NewClient("https://xtsujlxfyspxskmpsreg.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Inh0c3VqbHhmeXNweHNrbXBzcmVnIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTMyODUzNTQsImV4cCI6MjAwODg2MTM1NH0.-iN2Pgxhq9Q0bLj28bwKAT-PdulumbQ8lFBp66qNEQI", nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("tips").Select("*", "exact", false).Execute()

	fmt.Printf("%s, %d, %s", data, count, err)
}
