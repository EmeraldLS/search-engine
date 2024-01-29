package main

import (
	"flag"
	"log"
	"time"

	"github.com/EmeraldLS/search-engine/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "Wiki abstract dump path")
	flag.StringVar(&query, "q", "Who is wizkid?", "Query string")
	flag.Parse()

	log.Println("Full text search is in progress...")
	start := time.Now()

	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents took %s", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)

	idx.Add(docs...)
	log.Printf("Indexed %d documents took %s", len(docs), time.Since(start))

	start = time.Now()
	matchesIDs := idx.Search(query)
	log.Printf("Found %d matches took %s", len(matchesIDs), time.Since(start))

	for _, id := range matchesIDs {
		doc := docs[id]
		log.Printf("ID: %d, Title: %s\n", doc.ID, doc.Title)
	}

}
