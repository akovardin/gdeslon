package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/horechek/gdeslon"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	client := gdeslon.NewClient(
		"https://api.gdeslon.ru/api",
		"xxx",
	)

	type Response struct {
		Offers []struct {
			Id          uint64 `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Url         string `json:"url"`
		} `json:"offers"`
	}

	r := Response{}
	params := url.Values{}
	params.Add("q", "книги")
	err := client.Call("search.json", "GET", params, &r)
	if err != nil {
		log.Fatal(err)
	}

	for _, offer := range r.Offers {
		fmt.Println("==============")
		fmt.Println("id:", offer.Id)
		fmt.Println("name:", offer.Name)
		fmt.Println("description:", offer.Description)
		fmt.Println("url:", offer.Url)
	}
}
