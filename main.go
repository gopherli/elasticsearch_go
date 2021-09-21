// package main

// import (
// 	"elasticsearch_go/internal/dao/es"
// 	"elasticsearch_go/internal/routers"
// 	"net/http"
// 	"time"
// )

// func init() {
// 	es.InitClient()
// }

// func main() {
// 	router := routers.NewRouter()
// 	s := &http.Server{
// 		Addr:           ":8090",
// 		Handler:        router,
// 		ReadTimeout:    5 * time.Second,
// 		WriteTimeout:   5 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	s.ListenAndServe()
// }
// $ go run _examples/main.go

package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

func main() {
	log.SetFlags(0)

	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	)

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	//
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print version number.
	log.Printf("~~~~~~~> Elasticsearch %s", r["version"].(map[string]interface{})["number"])

	// 2. Index documents concurrently
	//
	for i, title := range []string{"Book Er", "Book Dan"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Set up the request object directly.
			req := esapi.IndexRequest{
				Index:      "book",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(`{"title" : "` + title + `"}`),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Use the helper methods of the client.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("book"),
		es.Search.WithBody(strings.NewReader(`{"query" : { "match" : { "title" : "book" } }}`)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}
