package main

import (
	"fmt"

	directus "github.com/marlaone/directus-go-sdk"
)

func main() {
	d := directus.NewDirectus(directus.NewConfigWithEndpoint("http://localhost:8055"))

	err := d.GetClient().Login("jps@marla.one", "2RZ9dfU*YV")

	if err != nil {
		panic(err)
	}

	collections, err := d.GetCollections()

	if err != nil {
		panic(err)
	}

	fmt.Println(collections)

	collection, err := d.GetCollection("directus_activity")

	if err != nil {
		panic(err)
	}

	fmt.Println(collection)

	q := directus.NewQuery()

	items, err := d.GetItems("test", q)

	if err != nil {
		panic(err)
	}

	fmt.Println(items)
}
