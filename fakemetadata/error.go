// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"net/http"

	"github.com/google/go-safeweb/safehttp"
)

// StatusError represents an error and safehttp.StatusCode.
//
// This error requires custom safehttp dispatcher.
type StatusError struct {
	err    error
	status safehttp.StatusCode
}

// NewStatusError returns the new StatusError from err and status args.
func NewStatusError(err error, status safehttp.StatusCode) StatusError {
	return StatusError{
		err:    err,
		status: status,
	}
}

// Code implements safehttp.ErrorResponse.Code.
func (e StatusError) Code() safehttp.StatusCode {
	return e.status
}

// Error implements safehttp.Dispatcher.Error.
func (e StatusError) Error(w http.ResponseWriter, resp safehttp.ErrorResponse) error {
	http.Error(w, e.err.Error(), int(e.Code()))
	return nil
}
