// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"net/http"

	"github.com/google/go-safeweb/safehttp"
)

// Dispatcher is a custom safehttp.Dispatcher implementation.
// See:
//
//	https://pkg.go.dev/github.com/google/go-safeweb/safehttp#hdr-Dispatcher.
type Dispatcher struct {
	// No need for a Write method, the default dispatcher knows how to write all
	// non-error responses we use in this project.
	safehttp.DefaultDispatcher
}

func (d Dispatcher) Error(rw http.ResponseWriter, resp safehttp.ErrorResponse) error {
	switch x := resp.(type) {
	case StatusError:
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.WriteHeader(int(x.Code()))
		return x
	}

	// calling the default dispatcher in case we have no custom responses that match.
	// This is strongly advised.
	return d.DefaultDispatcher.Error(rw, resp)
}
