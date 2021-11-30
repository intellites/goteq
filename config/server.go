package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/intellites/goteq/routers"
)

func NewServer() {
	// Init routers
	router := routers.NewRouter()

	// CORS
	router.Use(mux.CORSMethodMiddleware(router))

	log.Printf("Starting server on port %d", env.Port)

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", env.Host, env.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Server Shutting down")
	os.Exit(0)
}
