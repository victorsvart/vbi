package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/victorsvart/vbi/internal/wiring"
)

func main() {
	app := wiring.WireApp()
	sv := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	go func() {
		if err := sv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server start err := %v", err)
		}

		log.Println("Server simply stopped...")
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	ctx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdown()

	if err := sv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown err: %v", err)
	}

	log.Println("Server stopped. Bye bye~")
}
