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

	result := client.Rpc("add_them", "", map[string]int{"a": 12, "b": 3})

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	fmt.Println(result)
	data, count, err := client.From("tips").Select("*", "exact", false).Execute()

	fmt.Printf("%s, %d, %s", data, count, err)
}
