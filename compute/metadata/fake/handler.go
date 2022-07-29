// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"github.com/google/go-safeweb/safehttp"
)

type rootResponse struct{}

var _ safehttp.ErrorResponse = rootResponse{}

// Code implements safehttp.ErrorResponse.Code.
func (rootResponse) Code() safehttp.StatusCode {
	return safehttp.StatusOK
}

func rootHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	return w.WriteError(rootResponse{})
}
