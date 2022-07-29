// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"os"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
)

func rootHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	return w.WriteError(safehttp.StatusOK)
}

var projectEnvs = []string{"GOOGLE_CLOUD_PROJECT", "GCP_PROJECT", "GOOGLE_GCP_PROJECT"}

func projectHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range projectEnvs {
		if proj, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(proj))
		}
	}

	return w.WriteError(safehttp.StatusNotFound)
}

var numericProjectEnvs = []string{"GOOGLE_CLOUD_NUMERIC_PROJECT", "GCP_NUMERIC_PROJECT", "GOOGLE_GCP_NUMERIC_PROJECT"}

func numericProjectHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range numericProjectEnvs {
		if proj, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(proj))
		}
	}

	return w.WriteError(safehttp.StatusNotFound)
}

var hostnameEnvs = []string{"GOOGLE_HOSTNAME"}

func hostnameHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range hostnameEnvs {
		if hostname, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(hostname))
		}
	}

	return w.WriteError(safehttp.StatusNotFound)
}
