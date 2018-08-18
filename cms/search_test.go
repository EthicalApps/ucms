package cms_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethicalapps/ucms/cms"
	"github.com/ethicalapps/ucms/cms/store/bolt"
)

func TestSearch(t *testing.T) {
	db := "test.db"

	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	store, err := bolt.New(filepath.Join(dir, db))
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	cms.Init(dir, store)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := cms.NewRepository("repo")
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
