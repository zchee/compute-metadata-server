package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/zchee/compute-metadata-server/fakemetadata"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	srv := fakemetadata.NewServer()
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("MetadataHostEnv: %s\n", os.Getenv(fakemetadata.MetadataHostEnv))
	<-ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
