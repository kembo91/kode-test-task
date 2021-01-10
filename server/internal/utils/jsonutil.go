package utils

import (
	"encoding/json"
	"net/http"
)

//JSONErrorStruct is a struct for json errors
type JSONErrorStruct struct {
	Error string `json:"error"`
}

type jsonSuccessStruct struct {
	Result string `json:"result"`
}

//JSONError is a function used to return JSON error
func JSONError(w http.ResponseWriter, e error, status int) {
	var d JSONErrorStruct
	d.Error = e.Error()
	js, _ := json.Marshal(d)
	http.Error(w, string(js), status)
}

//JSONSuccess is a function for returning a generic success response
func JSONSuccess(w http.ResponseWriter) {
	var r jsonSuccessStruct
	r.Result = "success"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(r)
}
