// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"strings"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
)

var endpoints = []string{
	"instance/",
	"oslogin/",
	"project/",
}

func rootHandler(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
	switch r.URL().Path() {
	case "/", "/computeMetadata":
		return w.Write(safehtml.HTMLEscaped("computeMetadata/"))
	case "/computeMetadata/", "/computeMetadata/v1":
		return w.Write(safehtml.HTMLEscaped("v1/"))
	case "/computeMetadata/v1/":
		return w.Write(safehtml.HTMLEscaped(strings.Join(endpoints, "\n")))
	}

	return safehttp.NotWritten()
}

func redirectHandler(to string) safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.Redirect(w, r, to, safehttp.StatusFound)
	})
}
