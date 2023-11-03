package supabase

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

func Connect() {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("tips").Select("*", "exact", false).Execute()

	fmt.Printf("%s, %d, %s", data, count, err)
}
