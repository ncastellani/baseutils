package baseutils

import (
	"math/rand"
	"time"
)

// Empty
// define an empty interface literal type to pass null as a parameter
var Empty interface{}

// StringInSlice
// check if the informed string is on the passed slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

// IsMap
// check if the informed interface is an map of string to interface
func IsMap(v interface{}) bool {
	switch v.(type) {
	case map[string]interface{}:
		return true
	}

	return false
}

// RandomString
// generate a random string of the passed length using the desired chars
func RandomString(length int, upperCase, lowerCase, numbers bool) string {

	// prepare the charset
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if upperCase {
		charset = charset + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if lowerCase {
		charset = charset + "abcdefghijklmnopqrstuvwxyz"
	}

	if numbers {
		charset = charset + "0123456789"
	}

	// generate the random string and return
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano() - int64(rand.Int())))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

// GetKey
// get a key value of a interface within a string map
func GetKey(needle string, haystack map[string]interface{}) interface{} {
	if val, ok := haystack[needle]; ok {
		return val
	} else {
		return nil
	}
}
