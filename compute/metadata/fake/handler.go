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

const (
	EnvGoogleCloudProject = "GOOGLE_CLOUD_PROJECT"
	EnvGCPProject         = "GCP_PROJECT"
	EnvGoogleGCPProject   = "GOOGLE_GCP_PROJECT"
)

var projectEnvs = []string{EnvGoogleCloudProject, EnvGCPProject, EnvGoogleGCPProject}

func projectHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range projectEnvs {
		if proj, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(proj))
		}
	}

	return w.WriteError(safehttp.StatusNotFound)
}

const (
	EnvGoogleCloudNumericProject = "GOOGLE_CLOUD_NUMERIC_PROJECT"
	EnvGCPNumeriCProject         = "GCP_NUMERIC_PROJECT"
	EnvGoogleGCPNumericProject   = "GOOGLE_GCP_NUMERIC_PROJECT"
)

var numericProjectEnvs = []string{EnvGoogleCloudNumericProject, EnvGCPNumeriCProject, EnvGoogleGCPNumericProject}

func numericProjectHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	for _, env := range numericProjectEnvs {
		if proj, ok := os.LookupEnv(env); ok {
			return w.Write(safehtml.HTMLEscaped(proj))
		}
	}

	return w.WriteError(safehttp.StatusNotFound)
}

func redirectHandler(to string) safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.Redirect(w, r, to, safehttp.StatusFound)
	})
}
