package es

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// CreateIndex：创建索引
func (e EsOpreate) CreateIndex() (*esapi.Response, error) {
	response, err := client.Indices.Create("book-0.1.0", client.Indices.Create.WithBody(strings.NewReader(`
	{
		"aliases": {
			"book":{}
		},
		"settings": {
			"analysis": {
				"normalizer": {
					"lowercase": {
						"type": "custom",
						"char_filter": [],
						"filter": ["lowercase"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "keyword",
					"normalizer": "lowercase"
				},
				"price": {
					"type": "double"
				},
				"summary": {
					"type": "text",
					"analyzer": "ik_max_word"
				},
				"author": {
					"type": "keyword"
				},
				"pubDate": {
					"type": "date"
				},
				"pages": {
					"type": "integer"
				}
			}
		}
	}
	`)))
	if err != nil {
		return nil, err
	}
	return response, nil
}
