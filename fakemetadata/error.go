// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
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

// Error returns a string representation of the saErrorResponse.
func (e StatusError) Error() string {
	return e.err.Error()
}

// Code implements safehttp.ErrorResponse.Code.
func (e StatusError) Code() safehttp.StatusCode {
	return e.status
}
