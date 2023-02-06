package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// File contains two methods that decodes json input to a golang type, Both were added just to see the different ways
// it can be done. It is recommended to use json.decode if your payload is coming from an io.Reader(e.g http request body)
// and use unmarshall if you already have the JSON data in memory(like a variable).
/**
https://stackoverflow.com/questions/21197239/decoding-json-using-json-unmarshal-vs-json-newdecoder-decode
https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
*/

func DecodeJson(r *http.Request, x interface{}) error {
	return json.NewDecoder(r.Body).Decode(x)
}

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