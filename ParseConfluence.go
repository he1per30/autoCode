package parse

import (
	"autoCode/internal"
	"fmt"
	goconfluence "github.com/biome-search/confluence-go-api"
	"log"
)

func Parse() {
	api, err := goconfluence.NewAPI(internal.URL, internal.UserName, internal.TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(api)
}
