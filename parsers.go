package baseutils

import (
	"encoding/json"
	"os"
)

// ParseJSONFromFile
// parse a JSON file to an interface
func ParseJSON(path string, data interface{}) (err error) {

	// try to open the JSON file
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}

	// parse JSON file to interface
	return json.Unmarshal([]byte(file), data)
}
