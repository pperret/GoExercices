// Package params provides a reflection-based builder for URL parameters.
package params

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Pack create a URL from struct fields
func Pack(u *url.URL, ptr interface{}) error {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	if v.Type().Kind() != reflect.Struct {
		return fmt.Errorf("%v is not a struct", ptr)
	}

	q := u.Query()
	for i := 0; i < v.NumField(); i++ {

		// Compute parameter name
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		// Encode the field
		if err := encode(q, name, v.Field(i)); err != nil {
			return err
		}
	}
	u.RawQuery = q.Encode()
	return nil
}

// encode encodes one field as URL parameters
func encode(q url.Values, name string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err := encode(q, name, v.Index(i)); err != nil {
				return nil
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		q.Add(name, strconv.FormatInt(v.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		q.Add(name, strconv.FormatUint(v.Uint(), 10))
	case reflect.String:
		q.Add(name, v.String())
	default:
		return fmt.Errorf("invalid type %v", v.Kind())
	}
	return nil
}
