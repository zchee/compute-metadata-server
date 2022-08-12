// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata_test

import (
	"testing"

	"github.com/zchee/compute-metadata-server/fakemetadata"
)

func TestOnTest(t *testing.T) {
	if ok := fakemetadata.OnTest(); !ok {
		t.Fatal("expected true")
	}
}
