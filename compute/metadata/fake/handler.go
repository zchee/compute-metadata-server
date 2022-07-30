// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"github.com/google/go-safeweb/safehttp"
)

func rootHandler(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest) safehttp.Result {
	return w.WriteError(safehttp.StatusOK)
}

func redirectHandler(to string) safehttp.Handler {
	return safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
		return safehttp.Redirect(w, r, to, safehttp.StatusFound)
	})
}
