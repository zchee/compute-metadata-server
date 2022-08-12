// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/google/go-safeweb/safehttp"
)

// JSONResponse should encapsulate a valid JSON object that will be serialised
// and written to the http.ResponseWriter using a JSON encoder.
type JSONResponse struct {
	Data interface{}
}

// WriteJSON creates a JSONResponse from the data object and calls the Write
// function of the ResponseWriter, passing the response.
//
// The data object should be valid JSON, otherwise an error will occur.
func WriteJSON(w safehttp.ResponseWriter, data interface{}) safehttp.Result {
	return w.Write(JSONResponse{data})
}

// Dispatcher is a custom safehttp.Dispatcher implementation.
// See:
//
//	https://pkg.go.dev/github.com/google/go-safeweb/safehttp#hdr-Dispatcher.
type Dispatcher struct {
	safehttp.DefaultDispatcher
}

// Write implemens safehttp.Dispatcher.Write.
func (d Dispatcher) Write(rw http.ResponseWriter, resp safehttp.Response) error {
	switch x := resp.(type) {
	case JSONResponse:
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(rw).Encode(x.Data)
	}

	// calling the default dispatcher in case we have no custom responses that match.
	// This is strongly advised.
	return d.DefaultDispatcher.Write(rw, resp)
}

// Error implemens safehttp.Dispatcher.Error.
func (d Dispatcher) Error(rw http.ResponseWriter, resp safehttp.ErrorResponse) error {
	switch x := resp.(type) {
	case StatusError:
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		return x.Error(rw, resp)
	}

	// calling the default dispatcher in case we have no custom responses that match.
	// This is strongly advised.
	return d.DefaultDispatcher.Error(rw, resp)
}
