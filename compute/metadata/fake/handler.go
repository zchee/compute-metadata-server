// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	rootPattern = "/computeMetadata/v1"
)

const (
	hdrContentType = "Content-Type"
	mimeTextHTML   = "text/html; charset=UTF-8"
)

const (
	MetadataFlavorHeader = "Metadata-Flavor"
	MetadataFlavorValue  = "Google"
)

func checkHeader(h http.Header) error {
	metadataFlavor := h.Get(MetadataFlavorHeader)
	if !strings.EqualFold(metadataFlavor, MetadataFlavorValue) {
		return fmt.Errorf("%s header is wrong: %s", MetadataFlavorHeader, metadataFlavor)
	}

	return nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if err := checkHeader(r.Header); err != nil {
		w.Header().Set(hdrContentType, mimeTextHTML)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
