// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !purego
// +build !purego

// Package metadata provides access to Google Compute Engine (GCE)
// metadata and API service accounts.
//
// This package is a wrapper around the GCE metadata service,
// as documented at https://cloud.google.com/compute/docs/metadata/overview.
//
// This package holds "cloud.google.com/go/compute/metadata@v1.6.1" function and type aliases.
//
// TODO(zchee): auto generate.
package metadata

import (
	"net/http"
	_ "unsafe" // for go:linkname

	"cloud.google.com/go/compute/metadata"
)

// A Client provides metadata.
type Client = metadata.Client

// Error contains an error response from the server.
type Error = metadata.Error

// NotDefinedError is returned when requested metadata is not defined.
//
// The underlying string is the suffix after "/computeMetadata/v1/".
//
// This error is not returned if the value is defined to be the empty
// string.
type NotDefinedError = metadata.NotDefinedError

// Email calls Client.Email on the default client.
//
//go:linkname Email cloud.google.com/go/compute/metadata.Email
func Email(serviceAccount string) (string, error)

// ExternalIP returns the instance's primary external (public) IP address.
//
//go:linkname ExternalIP cloud.google.com/go/compute/metadata.ExternalIP
func ExternalIP() (string, error)

// Get calls Client.Get on the default client.
//
//go:linkname Get cloud.google.com/go/compute/metadata.Get
func Get(suffix string) (string, error)

// Hostname returns the instance's hostname. This will be of the form
// "<instanceID>.c.<projID>.internal".
//
//go:linkname Hostname cloud.google.com/go/compute/metadata.Hostname
func Hostname() (string, error)

// InstanceAttributeValue calls Client.InstanceAttributeValue on the default
// client.
//
//go:linkname InstanceAttributeValue cloud.google.com/go/compute/metadata.InstanceAttributeValue
func InstanceAttributeValue(attr string) (string, error)

// InstanceID returns the current VM's numeric instance ID.
//
//go:linkname InstanceID cloud.google.com/go/compute/metadata.InstanceID
func InstanceID() (string, error)

// InstanceName returns the current VM's instance ID string.
//
//go:linkname InstanceName cloud.google.com/go/compute/metadata.InstanceName
func InstanceName() (string, error)

// InstanceTags returns the list of user-defined instance tags, assigned when
// initially creating a GCE instance.
//
//go:linkname InstanceTags cloud.google.com/go/compute/metadata.InstanceTags
func InstanceTags() ([]string, error)

// InternalIP returns the instance's primary internal IP address.
//
//go:linkname InternalIP cloud.google.com/go/compute/metadata.InternalIP
func InternalIP() (string, error)

// NumericProjectID returns the current instance's numeric project ID.
//
//go:linkname NumericProjectID cloud.google.com/go/compute/metadata.NumericProjectID
func NumericProjectID() (string, error)

// OnGCE reports whether this process is running on Google Compute Engine.
//
//go:linkname OnGCE cloud.google.com/go/compute/metadata.OnGCE
func OnGCE() bool

// ProjectAttributeValue calls Client.ProjectAttributeValue on the default
// client.
//
//go:linkname ProjectAttributeValue cloud.google.com/go/compute/metadata.ProjectAttributeValue
func ProjectAttributeValue(attr string) (string, error)

// ProjectAttributes calls Client.ProjectAttributes on the default client.
//
//go:linkname ProjectAttributes cloud.google.com/go/compute/metadata.ProjectAttributes
func ProjectAttributes() ([]string, error)

// ProjectID returns the current instance's project ID string.
//
//go:linkname ProjectID cloud.google.com/go/compute/metadata.ProjectID
func ProjectID() (string, error)

// Scopes calls Client.Scopes on the default client.
//
//go:linkname Scopes cloud.google.com/go/compute/metadata.Scopes
func Scopes(serviceAccount string) ([]string, error)

// Subscribe calls Client.Subscribe on the default client.
//
//go:linkname Subscribe cloud.google.com/go/compute/metadata.Subscribe
func Subscribe(suffix string, fn func(v string, ok bool) error) error

// Zone returns the current VM's zone, such as "us-central1-b".
//
//go:linkname Zone cloud.google.com/go/compute/metadata.Zone
func Zone() (string, error)

// NewClient returns a Client that can be used to fetch metadata. Returns the
// client that uses the specified http.Client for HTTP requests. If nil is
// Scopes calls Client.Scopes on the default client.
//
//go:linkname NewClient cloud.google.com/go/compute/metadata.NewClient
func NewClient(c *http.Client) *Client
