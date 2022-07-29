// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"os"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
)

type rootResponse struct{}

var _ safehttp.ErrorResponse = rootResponse{}

// Code implements safehttp.ErrorResponse.Code.
func (rootResponse) Code() safehttp.StatusCode {
	return safehttp.StatusOK
}

type notDefinedResponse struct{}

var _ safehttp.ErrorResponse = notDefinedResponse{}

// Code implements safehttp.ErrorResponse.Code.
func (notDefinedResponse) Code() safehttp.StatusCode {
	return safehttp.StatusNotFound
}

func rootHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	return w.WriteError(rootResponse{})
}

var projectEnvs = []string{"GOOGLE_CLOUD_PROJECT", "GCP_PROJECT", "GOOGLE_GCP_PROJECT"}

func projectHandler(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range projectEnvs {
		if proj, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(proj))
		}
	}

	return w.WriteError(notDefinedResponse{})
}

var numericProjectEnvs = []string{"GOOGLE_CLOUD_NUMERIC_PROJECT", "GCP_NUMERIC_PROJECT", "GOOGLE_GCP_NUMERIC_PROJECT"}

func numericProjectHandler(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range numericProjectEnvs {
		if proj, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(proj))
		}
	}

	return w.WriteError(notDefinedResponse{})
}
