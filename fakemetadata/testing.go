// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"os"
	"strings"
)

// OnTest reports whether the current state is on test.
func OnTest() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}
