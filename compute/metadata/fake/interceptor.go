// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"fmt"
	"strings"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
)

const (
	MetadataFlavorHeader = "Metadata-Flavor"
	MetadataFlavorValue  = "Google"

	ServerHeader = "Server"
	ServerValue  = "Metadata Server for VM"
)

type metadataInterceptor struct{}

var _ safehttp.Interceptor = (*metadataInterceptor)(nil)

// Before implements safehttp.Interceptor.Before.
func (metadataInterceptor) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, cfg safehttp.InterceptorConfig) safehttp.Result {
	metadataFlavor := r.Header.Get(MetadataFlavorHeader)
	if !strings.EqualFold(metadataFlavor, MetadataFlavorValue) {
		return w.Write(safehtml.HTMLEscaped(fmt.Sprintf("%s header is wrong: %s", MetadataFlavorHeader, metadataFlavor)))
	}

	return safehttp.NotWritten()
}

// Commit implements safehttp.Interceptor.Commit.
func (metadataInterceptor) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	// nothing to do
}

// Match implements safehttp.Interceptor.Match.
func (metadataInterceptor) Match(safehttp.InterceptorConfig) bool {
	return false
}

type serverInterceptor struct{}

var _ safehttp.Interceptor = (*serverInterceptor)(nil)

// Before implements safehttp.Interceptor.Before.
func (serverInterceptor) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, cfg safehttp.InterceptorConfig) safehttp.Result {
	setServer := w.Header().Claim(ServerHeader)
	setServer([]string{ServerValue})

	return safehttp.NotWritten()
}

// Commit implements safehttp.Interceptor.Commit.
func (serverInterceptor) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	// nothing to do
}

// Match implements safehttp.Interceptor.Match.
func (serverInterceptor) Match(safehttp.InterceptorConfig) bool {
	return false
}
