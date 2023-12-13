package main

import (
	"errors"
	"fmt"
	"github.com/xuche123/cloudgo/config"
	"io"
	"log"
	"net/http"
)

func main() {
	c := config.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server on :8080")
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
