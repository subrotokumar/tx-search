package main

import (
	"flag"
	"log"
	"time"

	"github.com/subrotokumar/tx-search/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "data/enwiki-latest-abstract.xml.gz", "Data Dump Path")
	flag.StringVar(&query, "q", "Small wild cat", "Search Query")
	flag.Parse()
	log.Println("Text Search is in progress")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d document is %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
