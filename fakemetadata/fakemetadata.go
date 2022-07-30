// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

// Package fakemetadata provides the fake GCE compute metadata server for testing.
package fakemetadata

import (
	"net/url"
	"path"
	"strings"
)

// List of metadata server variables.
//
// Those are exported and don't use constant types so can be replaced.
//
// See details: https://cloud.google.com/compute/docs/metadata/overview#parts-of-a-request
var (
	// metadataIP is the documented metadata server IP address.
	MetadataIP = "169.254.169.254"

	// MetadataHostEnv is the environment variable specifying the GCE metadata hostname.
	// If empty, the default value of metadataIP ("169.254.169.254") is used instead.
	//
	// The cloud.google.com/go/compute/metadata package maintainer said:
	// > This is variable name is not defined by any spec, as far as I know; it was made up for the Go package.
	//
	// So this environment variable is helpful to replace the server that the cloud.google.com/go/compute/metadata package accesses during testing.
	MetadataHostEnv = "GCE_METADATA_HOST"

	// RootURL is the documented metadata server Host.
	RootURL = "metadata.google.internal"

	// SubPath is the documented metadata sub-path.
	SubPath = path.Join("computeMetadata", "v1")
)

var rootURL = &url.URL{
	Scheme: "http",
	Host:   RootURL,
	Path:   SubPath,
}

// List of request http header constants.
//
// See also: https://cloud.google.com/compute/docs/metadata/overview
const (
	// RequestHeader is the required http header for access to the metadata server.
	//
	// This header indicates that the request was sent with the intention of retrieving metadata values,
	// rather than unintentionally from an insecure source, and lets the metadata server return the data you requested.
	// If you don't provide this header, the metadata server denies your request.
	RequestHeader = "Metadata-Flavor: Google"

	// LegacyRequestHeader is the legacy (but still supported) required http header for access to the metadata server.
	LegacyRequestHeader = "X-Google-Metadata-Request: True"
)

// QueryReplacer replaces string pairs for not supported in a request path to the metadata server.
//
// See details: https://cloud.google.com/compute/docs/metadata/overview#limitations
var QueryReplacer = strings.NewReplacer(
	"%21", "!",
	"%24", "$",
	"%27", "'",
	"%28", "(",
	"%29", ")",
	"%2A", "*",
	"%2C", ",",
	"%40", "@",
)
