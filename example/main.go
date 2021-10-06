package main

import (
	"fmt"

	directus "github.com/marlaone/directus-go-sdk"
)

func main() {
	directus := directus.NewDirectus(directus.NewConfigWithEndpoint("http://localhost:8055"))

	err := directus.GetClient().Login("jps@marla.one", "2RZ9dfU*YV")

	if err != nil {
		panic(err)
	}

	collections, err := directus.GetCollections()

	if err != nil {
		panic(err)
	}

	fmt.Println(collections)
}
