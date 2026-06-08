package v1

import (
	"encoding/json"
	"net/http"
)

const internalServerError string = `{"status":"ERROR", "description":"Internal server error"}`
const okResult string = `{"status":"OK"}`
const badRequest string = `{"status":"ERROR", "description":"Bad request"}`

type responseStatus string

const (
	responseStatusOk    = "OK"
	responseStatusError = "ERROR"
)

type SimpleResponse struct {
	Status      responseStatus `json:"status"`
	Description string         `json:"description"`
}

func newSimpleResponse(status responseStatus, descr string) SimpleResponse {
	return SimpleResponse{
		Status:      status,
		Description: descr,
	}
}

func makeAnyError(w http.ResponseWriter, err error, statusCode http.ConnState) {
	errBody := newSimpleResponse(responseStatusError, err.Error())
	jsonedBody, err := json.Marshal(errBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalServerError))
	} else {
		w.WriteHeader(int(statusCode))
		w.Write(jsonedBody)

	}
}

func makeBadRequestError(w http.ResponseWriter, err error) {
	makeAnyError(w, err, http.StatusBadRequest)
}

func makeInternalServerError(w http.ResponseWriter, err error) {
	makeAnyError(w, err, http.StatusInternalServerError)
}
