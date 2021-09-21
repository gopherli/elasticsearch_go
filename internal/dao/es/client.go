package es

import (
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
)

type EsOpreate struct{}

var (
	client *es.Client
)

func ReturnNewEsOpreate() EsOpreate {
	return EsOpreate{}
}

func InitClient() {
	var err error
	client, err = es.NewClient(es.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "username",
		Password:  "password",
	})
	if err != nil {
		log.Fatal(err)
	}
}
