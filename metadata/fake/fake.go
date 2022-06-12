// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	_ "cloud.google.com/go/iam/credentials/apiv1"
	_ "golang.org/x/net/http2"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/oauth2/google"
	_ "google.golang.org/api/idtoken"
	_ "google.golang.org/api/impersonate"
	_ "google.golang.org/genproto/googleapis/iam/credentials/v1"
)
