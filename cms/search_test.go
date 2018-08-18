package cms_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ethicalapps/ucms/cms"
	"github.com/ethicalapps/ucms/cms/bolt"
)

func TestSearch(t *testing.T) {
	db := "search_test.db"

	store, err := bolt.New(db)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	cms.Init(store)

	repo, err := cms.NewRepository("search_test")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile("testdata/types/article.json")
	if err != nil {
		t.Error("ERROR:", err)
	}
	if err := repo.PutType("article", content); err != nil {
		t.Error("ERROR:", err)
	}
	content, err = ioutil.ReadFile("testdata/documents/article_1.json")
	if err != nil {
		t.Error("ERROR:", err)
	}
	if err := repo.PutDocument("article", "article1", content); err != nil {
		t.Error("ERROR:", err)
	}

	content, err = ioutil.ReadFile("testdata/documents/article_2.json")
	if err != nil {
		t.Error("ERROR:", err)
	}
	if err := repo.PutDocument("article", "article2", content); err != nil {
		t.Error("ERROR:", err)
	}

	results, err := repo.Query("article", "loverly", 10, 0)
	if err != nil {
		t.Error("ERROR:", err)
	}
	if len(results) != 1 {
		t.Error("Expected 1 result, got", len(results))
	}

	results, err = repo.Query("article", "article", 10, 0)
	if err != nil {
		t.Error("ERROR:", err)
	}
	if len(results) != 2 {
		t.Error("Expected 2 results, got", len(results))
	}

}
