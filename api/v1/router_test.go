package v1_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ethicalapps/ucms/api/v1"
	"github.com/ethicalapps/ucms/cms"
	"github.com/ethicalapps/ucms/cms/bolt"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	db := "test.db"

	store, err := bolt.New(db)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	cms.Init(store)

	server := httptest.NewServer(v1.Router())
	defer server.Close()

	if err := apiRequest(t, server, "/repo/schema/type.json", http.StatusOK); err != nil {
		t.Error(err)
	}

	store.Close()

	if err := os.Remove(db); err != nil {
		t.Error("ERROR:", err)
	}
}

func apiRequest(t *testing.T, server *httptest.Server, path string, expect int) error {
	resp, err := http.Get(server.URL + path)
	if err != nil {
		return err
	}

	assert.Equal(t, expect, resp.StatusCode)

	return nil
}
