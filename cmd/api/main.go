package main

import (
	"errors"
	"fmt"
	"github.com/xuche123/cloudgo/cmd/api/router"
	"github.com/xuche123/cloudgo/config"
	"log"
	"net/http"
)

//  @title          CLOUDGO API
//  @version        1.0
//  @description    This is a sample RESTful API with a CRUD

//  @contact.name   Xu Che
//  @contact.url    https://github.com/xuche123/cloudgo

//  @license.name   MIT License
//  @license.url    https://github.com/xuche123/cloudgo/blob/master/LICENSE

// @host       localhost:8080
// @basePath   /v1
func main() {
	c := config.New()
	r := router.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server on :8080")
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}
