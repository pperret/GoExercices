package sexpr

import (
	"reflect"
	"strings"
)

// tagDecode decodes the value of struct tag
func tagDecode(tag reflect.StructTag) (string, bool) {
	omitempty := false
	name := ""
	tagValue := tag.Get("sexpr")
	if tagValue != "" {
		tagValues := strings.Split(tagValue, ",")
		for _, o := range tagValues {
			v := strings.TrimSpace(o)
			if v == "omitempty" {
				omitempty = true
			} else {
				name = v
			}
		}
	}
	return name, omitempty
}
