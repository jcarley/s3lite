package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getHeaderValue(name string, req *http.Request) string {
	headerValue := req.Header[name]
	if len(headerValue) > 0 {
		return headerValue[0]
	}
	return ""
}

func encode(rw http.ResponseWriter, v interface{}) error {
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(v)
	if err != nil {
		return err
	}
	return nil
}

func decode(req *http.Request, v interface{}) error {
	// read raw bytes
	dataBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	// unmarsharl to map
	var data map[string]interface{}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return err
	}

	// decode to a struct
	err = mapstructure.Decode(data, v)
	if err != nil {
		return err
	}

	return nil
}

func httpError(err error, status int, rw http.ResponseWriter) {
	errorMessage := Message{"error", err.Error()}

	result, err := json.Marshal(errorMessage)
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	} else {
		http.Error(rw, string(result), status)
	}
}
