package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Convert incoming request payload from json to a datatype golang can use
// Method takes in request object and an "any" type of request body
func ConvertFromJson(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// All good unmarshal
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}