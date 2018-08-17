package cms_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ethicalapps/ucms/cms"
	"github.com/ethicalapps/ucms/cms/bolt"
)

func TestRepository(t *testing.T) {
	db := "test.db"

	store, err := bolt.New(db)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	cms.Init(store)

	repo, err := cms.NewRepository("blog")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile("testdata/types/person.json")
	if err != nil {
		t.Error("ERROR:", err)
	}
	if err := repo.PutType("person", content); err != nil {
		t.Error("ERROR:", err)
	}

	content, err = ioutil.ReadFile("testdata/types/article.json")
	if err != nil {
		t.Error("ERROR:", err)
	}
	if err := repo.PutType("article", content); err != nil {
		t.Error("ERROR:", err)
	}

	content, err = ioutil.ReadFile("testdata/documents/person_1.json")
	if err != nil {
		t.Error("ERROR:", err)
	}
	if err := repo.PutDocument("person", "person1", content); err != nil {
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

	docs, err := repo.ListDocuments("article")
	if len(docs) != 2 {
		t.Error("Expected 2 docs, got", len(docs))
	}

	store.Close()

	if err := os.Remove(db); err != nil {
		t.Error("ERROR:", err)
	}
}
