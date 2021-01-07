package utils

import (
	"encoding/json"
	"net/http"
)

//JSONErrorStruct is a struct for json errors
type JSONErrorStruct struct {
	Error string `json:"error"`
}

//JSONError is a function used to return JSON error
func JSONError(w http.ResponseWriter, e error, status int) {
	var d JSONErrorStruct
	d.Error = e.Error()
	js, _ := json.Marshal(d)
	http.Error(w, string(js), status)
}
