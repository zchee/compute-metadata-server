// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var portRandSrc = rand.NewSource(time.Now().UTC().UnixNano())

// randomPort generates random port number and checks whether the it available(unused) port.
func randomPort(network string) string {
	var p string
	rnd := rand.New(portRandSrc)
	for {
		p = strconv.Itoa(rnd.Intn(55535) + 1000)
		if conn, err := net.Dial(network, p); err == nil {
			conn.Close()
			continue
		}
		return p
	}
}

type Server struct {
	port string
	mux  *http.ServeMux
}

// NewServer returns the new fake metadata server.
func NewServer() *Server {
	port := randomPort("tcp4")

	// inject MetadataHostEnv hostname
	os.Setenv(MetadataHostEnv, net.JoinHostPort("localhost", port))

	mux := http.NewServeMux()
	mux.HandleFunc(rootPattern, rootHandler)

	return &Server{
		port: port,
		mux:  mux,
	}
}
