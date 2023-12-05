package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIResult struct {
	Result struct {
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Errors       string `json:"errors"`
	} `json:"result"`
	Data interface{} `json:"data"`
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	var payload APIResult
	payload.Data = data

	out, err := json.Marshal(payload)

	if err != nil {
		log.Println("Error Marshalling: ", err)
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		log.Println("Error writing to http:", err)
		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, errorMessage string, errorCode string, status ...int) error {

	log.Println("Internal Error:", err)

	statusCode := http.StatusInternalServerError

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload APIResult
	payload.Result.ErrorCode = errorCode
	payload.Result.ErrorMessage = errorMessage
	payload.Result.Errors = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	out, err := json.Marshal(payload)

	if err != nil {
		log.Println("Error Marshalling: ", err)
		return err
	}

	_, err = w.Write(out)

	if err != nil {
		log.Println("Error writing to http:", err)
		return err
	}

	return nil
}
