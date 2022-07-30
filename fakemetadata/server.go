// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/google/go-safeweb/safehttp"
	"golang.org/x/net/http2"
)

// Server represents a fake metadata server.
type Server struct {
	srv *safehttp.Server
}

// NewServer returns the new fake metadata server.
func NewServer() *Server {
	addr := net.JoinHostPort("localhost", randomPort("tcp4"))

	// inject MetadataHostEnv host
	os.Setenv(MetadataHostEnv, addr)

	muxConfig := safehttp.NewServeMuxConfig(nil)
	muxConfig.Intercept(metadataFlavorInterceptor{})
	muxConfig.Intercept(serverInterceptor{})
	muxConfig.Intercept(staticHeadersInterceptor{})

	mux := muxConfig.Mux()
	mux.Handle("/", safehttp.MethodGet, safehttp.HandlerFunc(rootHandler))
	mux.Handle("/computeMetadata", safehttp.MethodGet, safehttp.HandlerFunc(rootHandler))
	mux.Handle("/computeMetadata/", safehttp.MethodGet, safehttp.HandlerFunc(rootHandler))
	mux.Handle("/computeMetadata/v1", safehttp.MethodGet, safehttp.HandlerFunc(rootHandler))
	mux.Handle("/computeMetadata/v1/", safehttp.MethodGet, safehttp.HandlerFunc(rootHandler))
	(&ProjectHandler{}).RegisterHandlers(mux)
	(&InstanceHandler{}).RegisterHandlers(mux)

	return &Server{
		srv: &safehttp.Server{
			Addr: addr,
			Mux:  mux,
		},
	}
}

var portRandSrc = rand.NewSource(time.Now().UTC().UnixNano())

// randomPort generates random port number and checks whether the it available (unused) port.
func randomPort(network string) string {
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

// Addr returns the fake metadata server addr.
func (s *Server) Addr() string { return s.srv.Addr }

func buildStd(s *safehttp.Server) error {
	v := reflect.ValueOf(s).Elem()
	tartedVal := v.FieldByName("started")
	startedPtr := unsafe.Pointer(tartedVal.UnsafeAddr())

	if *(*bool)(startedPtr) {
		return errors.New("server already started")
	}

	srvVal := v.FieldByName("srv")
	srvPtr := unsafe.Pointer(srvVal.UnsafeAddr())
	if (*http.Server)(srvPtr) != nil {
		// Server was already built
		return nil
	}

	if s.Mux == nil {
		return errors.New("building server without a mux")
	}

	srv := &http.Server{
		Addr:           s.Addr,
		Handler:        s.Mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 10 * 1024,
	}
	if s.ReadTimeout != 0 {
		srv.ReadTimeout = s.ReadTimeout
	}
	if s.WriteTimeout != 0 {
		srv.WriteTimeout = s.WriteTimeout
	}
	if s.IdleTimeout != 0 {
		srv.IdleTimeout = s.IdleTimeout
	}
	if s.MaxHeaderBytes != 0 {
		srv.MaxHeaderBytes = s.MaxHeaderBytes
	}
	if s.TLSConfig != nil {
		cfg := s.TLSConfig.Clone()
		cfg.MinVersion = tls.VersionTLS12
		cfg.PreferServerCipherSuites = true
		srv.TLSConfig = cfg
	}
	for _, f := range s.OnShudown {
		srv.RegisterOnShutdown(f)
	}
	if s.DisableKeepAlives {
		srv.SetKeepAlivesEnabled(false)
	}

	http2.ConfigureServer(srv, &http2.Server{})

	*(*http.Server)(srvPtr) = *srv

	return nil
}

// ListenAndServe is a wrapper for https://pkg.go.dev/pkg/net/http/#Server.ListenAndServe
func (s *Server) ListenAndServe() error {
	if err := buildStd(s.srv); err != nil {
		return err
	}

	return s.srv.ListenAndServe()
}

// ListenAndServeTLS is a wrapper for https://pkg.go.dev/pkg/net/http/#Server.ListenAndServeTLS
func (s *Server) ListenAndServeTLS(certFile, keyFile string) error {
	if err := buildStd(s.srv); err != nil {
		return err
	}

	return s.srv.ListenAndServeTLS(certFile, keyFile)
}

// Serve is a wrapper for https://pkg.go.dev/pkg/net/http/#Server.Serve
func (s *Server) Serve(l net.Listener) error {
	if err := buildStd(s.srv); err != nil {
		return err
	}

	return s.srv.Serve(l)
}

// ServeTLS is a wrapper for https://pkg.go.dev/pkg/net/http/#Server.ServeTLS
func (s *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
	if err := buildStd(s.srv); err != nil {
		return err
	}

	return s.srv.ServeTLS(l, certFile, keyFile)
}

// Shutdown is a wrapper for https://pkg.go.dev/pkg/net/http/#Server.Shutdown
func (s *Server) Shutdown(ctx context.Context) error {
	defer os.Unsetenv(MetadataHostEnv)

	return s.srv.Shutdown(ctx)
}

// Close is a wrapper for https://pkg.go.dev/pkg/net/http/#Server.Close
func (s *Server) Close() error {
	defer os.Unsetenv(MetadataHostEnv)

	return s.srv.Close()
}

var server unsafe.Pointer // *Server

// StartServer starts fake metadata server.
func StartServer() {
	srv := NewServer()
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	atomic.StorePointer(&server, unsafe.Pointer(srv))
}

// IsRunning reports whether the fake metadata server running.
func IsRunning() bool {
	return atomic.LoadPointer(&server) != nil
}

// Shutdown gracefully shuts down the fake metadata server.
func Shutdown(ctx context.Context) error {
	return (*Server)(atomic.LoadPointer(&server)).Shutdown(ctx)
}
