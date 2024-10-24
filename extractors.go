package baseutils

import (
	"gopkg.in/guregu/null.v4"
)

// ExtractNullInt
// extract a null.Int out of a interface that might not be int64
func ExtractNullInt(data interface{}) (v null.Int) {
	switch data.(type) {
	case float64:
		v = null.NewInt(int64(data.(float64)), true)
	}
	return
}

// ExtractNullString
// extract a null.String out of a interface that might not be string
func ExtractNullString(data interface{}) (v null.String) {
	switch data.(type) {
	case string:
		if data.(string) != "" {
			return null.NewString(data.(string), true)
		}
	}
	return
}

// ExtractStringArray
// extract a array of strings out of an interface
func ExtractStringArray(data interface{}) (v []string) {
	switch data.(type) {
	case []interface{}:
		for _, e := range data.([]interface{}) {
			switch e.(type) {
			case string:
				v = append(v, e.(string))
			}
		}
	}
	return
}
