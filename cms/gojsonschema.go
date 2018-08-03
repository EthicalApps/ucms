package cms

import (
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func init() {
	gojsonschema.FormatCheckers.Add("file", FileFormatChecker{})
}

// FileFormatChecker checks for file format
type FileFormatChecker struct{}

// IsFormat implements the gojsonschema.FormatChecker interface
func (f FileFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if ok == false {
		return false
	}

	return strings.HasPrefix(asString, "http")
}
