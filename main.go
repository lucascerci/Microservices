package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":3000", // configure the bind address
		Handler:      sm,      // set the default handler
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second, // max time to read request from the client
		WriteTimeout: 1 * time.Second, // max time to write response to the client
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	//Shutdown it will wait ultil the requests that are actually being handle by the server finish, then shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
