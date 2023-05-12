package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.3.74:9200"},
	}
	_, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	fmt.Println("Connect es successfully!!")
}
