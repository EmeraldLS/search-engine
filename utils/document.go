package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type Document struct {
	ID    int
	Title string `xml:"title"`
	Text  string `xml:"text"`
	URL   string `xml:"url"`
}

func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return []Document{}, err
	}

	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return []Document{}, err
	}

	defer gz.Close()

	doc := xml.NewDecoder(gz)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err = doc.Decode(&dump); err != nil {
		return []Document{}, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
