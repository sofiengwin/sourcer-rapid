package supabase

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

func main() {
	client, err := supabase.NewClient("API_URL", "API_KEY", nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()

	fmt.Printf("%s, %d, %s", data, count, err)
}

func Connect() {
	client, err := supabase.NewClient("API_URL", "API_KEY", nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("tips").Select("*", "exact", false).Execute()

	fmt.Printf("%s, %d, %s", data, count, err)
}
