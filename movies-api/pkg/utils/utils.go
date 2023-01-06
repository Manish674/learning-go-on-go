package utils

//code to umarshal data basically convert json to go

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// function to parse the json object into data that go understands
func parseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
