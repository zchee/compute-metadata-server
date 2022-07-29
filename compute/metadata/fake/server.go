// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/go-safeweb/safehttp/plugins/staticheaders"
)

var portRandSrc = rand.NewSource(time.Now().UTC().UnixNano())

// RandomPort generates random port number and checks whether the it available (unused) port.
func RandomPort(network string) string {
	var p string
	rnd := rand.New(portRandSrc)
	for {
		p = strconv.Itoa(rnd.Intn(55535) + 1000)
		if conn, err := net.Dial(network, p); err == nil {
			// not available
			conn.Close()
			continue
		}

		return p
	}
}

type Server struct {
	*safehttp.Server

	port string
}

// NewServer returns the new fake metadata server.
func NewServer() *Server {
	port := RandomPort("tcp4")
	addr := net.JoinHostPort("localhost", port)

	// inject MetadataHostEnv hostname
	os.Setenv(MetadataHostEnv, addr)

	muxConfig := safehttp.NewServeMuxConfig(nil)
	muxConfig.Intercept(metadataInterceptor{})
	muxConfig.Intercept(serverInterceptor{})
	muxConfig.Intercept(staticheaders.Interceptor{})

	mux := muxConfig.Mux()
	mux.Handle("/computeMetadata/v1", safehttp.MethodGet, safehttp.HandlerFunc(rootHandler))

	return &Server{
		Server: &safehttp.Server{
			Addr: addr,
			Mux:  mux,
		},
		port: port,
	}
}
