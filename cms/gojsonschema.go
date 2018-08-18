package cms

import (
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func init() {
	gojsonschema.FormatCheckers.Add("file", fileFormatChecker{})
}

// FileFormatChecker checks for file format
type fileFormatChecker struct{}

// IsFormat implements the gojsonschema.FormatChecker interface
func (f fileFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if ok == false {
		return false
	}

	return strings.HasPrefix(asString, "http")
}
