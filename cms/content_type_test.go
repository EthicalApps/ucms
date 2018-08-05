package cms_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/ethicalapps/ucms/cms"
)

func TestGenerateSchema(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/types/article.json")
	if err != nil {
		t.Error("There was an error:", err)
	}

	contentType := &cms.ContentType{}
	if err := json.Unmarshal(content, &contentType); err != nil {
		t.Error("There was an error:", err)
	}

	_, err = cms.GenerateSchema(contentType)
	if err != nil {
		t.Error("There was an error:", err)
	}
}
