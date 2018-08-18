package cms

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ContentType defines the structure
type ContentType struct {
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	DisplayField string   `json:"displayField"`
	Fields       []*Field `json:"fields"`
}

// Field defines the structure
type Field struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Items    struct {
		Type          string `json:"type"`
		ReferenceType string `json:"referenceType"`
	} `json:"items,omitempty"`
}

// GenerateSchema generates a JSON schema for a given ContentType
func GenerateSchema(c *ContentType) ([]byte, error) {
	required := []string{}
	properties := []string{}

	for _, field := range c.Fields {
		if field.Required {
			required = append(required, "\""+field.ID+"\"")
		}
		properties = append(properties, generateField(field))
	}

	schema := fmt.Sprintf(`
{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"id": "/schema/%s.json",
	"description": "%s",
	"type": "object",
	"properties": {%s},
	"additionalProperties": false,
	"required": [%s]
}`, c.ID, c.Description, strings.Join(properties, ","), strings.Join(required, ","))

	var clean map[string]interface{}
	if err := json.Unmarshal([]byte(schema), &clean); err != nil {
		return nil, err
	}

	out, err := json.MarshalIndent(clean, "", "    ")
	if err != nil {
		return nil, err
	}

	return out, nil
}

// GenerateIndexMapping generates a bleve search index mapping for a given
// ContentType
// func GenerateIndexMapping(c *ContentType) ([]byte, error) {
// }

func generateField(field *Field) string {
	return fmt.Sprintf(`
"%s": {
	"type": "%s"
}`, field.ID, schemaTypemap(field.Type))
}

func schemaTypemap(t string) string {
	if t == "Text" {
		return "string"
	} else if t == "TextArea" {
		return "string"
	} else if t == "Array" {
		return "array"
	} else if t == "Email" {
		return "string"
	}
	return "unknown"
}

// ContentTypeSchema is the core schema for all content types
const ContentTypeSchema = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "/schema/type.json",
  "type": "object",
  "properties": {
    "id": {
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
    "description": {
      "type": "string"
    },
    "displayField": {
      "type": "string"
    },
    "fields": {
      "type": "array",
      "items": [
        {
          "$ref": "#/definitions/field"
        }
      ]
    }
  },
  "additionalProperties": false,
  "required": [
    "id",
    "type",
    "name",
    "description",
    "displayField"
  ],
  "definitions": {
    "field": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "required": {
          "type": "boolean"
        },
        "items": {
          "type": "object",
          "properties": {
            "type": {
              "type": "string"
            },
            "referenceType": {
              "type": "string"
            }
          },
          "additionalProperties": false
        }
      },
      "additionalProperties": false,
      "required": [
        "id",
        "type",
        "required"
      ]
    }
  }
}`
