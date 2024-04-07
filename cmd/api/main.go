package main

import (
	"fmt"
	"helloworldapp/api/router"
	"helloworldapp/config"
	"log"
	"net/http"
)

//  @title          APP API
//  @version        1.0
//  @description    This is a sample RESTful API with a CRUD
//  @host       localhost:8080
//  @basePath   /v1
func main() {
    c := config.New()
    r := router.New()

    server := &http.Server{
        Addr: fmt.Sprintf(":%d", c.Server.Port),
        Handler: r,
        ReadTimeout: c.Server.TimeoutRead,
        WriteTimeout: c.Server.TimeoutWrite,
        IdleTimeout: c.Server.TimeoutIdle,
    }
    log.Println("Starting server on " + server.Addr)

    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("could not start server: %v", err)
    }
}

