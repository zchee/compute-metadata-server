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
)

type metadataFlavorInterceptor struct{}

var _ safehttp.Interceptor = metadataFlavorInterceptor{}

// Before implements safehttp.Interceptor.Before.
func (metadataFlavorInterceptor) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, cfg safehttp.InterceptorConfig) safehttp.Result {
	metadataFlavor := r.Header.Get(MetadataFlavorHeader)
	if !strings.EqualFold(metadataFlavor, MetadataFlavorValue) {
		return w.Write(safehtml.HTMLEscaped(fmt.Sprintf("%s header is wrong: %s", MetadataFlavorHeader, metadataFlavor)))
	}

	return safehttp.NotWritten()
}

// Commit claims and sets the following headers:
//  - Metadata-Flavor: Google
func (metadataFlavorInterceptor) Commit(w safehttp.ResponseHeadersWriter, r *safehttp.IncomingRequest, _ safehttp.Response, _ safehttp.InterceptorConfig) {
	h := w.Header()
	setMetadataFlavor := h.Claim(MetadataFlavorHeader)
	setMetadataFlavor([]string{MetadataFlavorValue})
}

// Match returns false since there are no supported configurations.
func (metadataFlavorInterceptor) Match(safehttp.InterceptorConfig) bool {
	return false
}

const (
	ServerHeader = "Server"
	ServerValue  = "Metadata Server for VM"
)

type serverInterceptor struct{}

var _ safehttp.Interceptor = serverInterceptor{}

// Before claims and sets the following headers:
//  - Server: Metadata Server for VM
func (serverInterceptor) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, _ safehttp.InterceptorConfig) safehttp.Result {
	setServer := w.Header().Claim(ServerHeader)
	setServer([]string{ServerValue})

	return safehttp.NotWritten()
}

// Commit is a no-op, required to satisfy the safehttp.Interceptor interface.
func (serverInterceptor) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	// nothing to do
}

// Match returns false since there are no supported configurations.
func (serverInterceptor) Match(safehttp.InterceptorConfig) bool {
	return false
}

const (
	XXSSProtectionHeader = "X-XSS-Protection"
	XXSSProtectionValue  = "0"

	XFrameOptionsHeader = "X-Frame-Options"
	XFrameOptionsValue  = "SAMEORIGIN"
)

// staticHeadersInterceptor claims and sets static headers on responses.
// The zero value is valid and ready to use.
type staticHeadersInterceptor struct{}

var _ safehttp.Interceptor = staticHeadersInterceptor{}

// Before claims and sets the following headers:
//  - X-XSS-Protection: 0
//  - X-Frame-Options: SAMEORIGIN
func (staticHeadersInterceptor) Before(w safehttp.ResponseWriter, r *safehttp.IncomingRequest, _ safehttp.InterceptorConfig) safehttp.Result {
	h := w.Header()
	setXXP := h.Claim(XXSSProtectionHeader)
	setXFO := h.Claim(XFrameOptionsHeader)
	setXXP([]string{XXSSProtectionValue})
	setXFO([]string{XFrameOptionsValue})

	return safehttp.NotWritten()
}

// Commit is a no-op, required to satisfy the safehttp.Interceptor interface.
func (staticHeadersInterceptor) Commit(safehttp.ResponseHeadersWriter, *safehttp.IncomingRequest, safehttp.Response, safehttp.InterceptorConfig) {
	// nothing to do
}

// Match returns false since there are no supported configurations.
func (staticHeadersInterceptor) Match(safehttp.InterceptorConfig) bool {
	return false
}
