package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/zchee/compute-metadata-server/fakemetadata"
)

var flagPort string

func main() {
	flag.StringVar(&flagPort, "port", "", "server port")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	srv := fakemetadata.NewServerWithPort(flagPort)
	errc := make(chan error, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errc <- err
		}
	}()

	fmt.Printf("MetadataHostEnv: %s\n", os.Getenv(fakemetadata.MetadataHostEnv))
	select {
	case err := <-errc:
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case <-ctx.Done():
		// nothing to do, initiate a graceful shutdown of the server
	}

	close(errc)
	cancel()
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
