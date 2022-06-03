package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/capslock-inc/gprc-demo/Api/handlers"
)

func main() {
	// logger
	logs := log.New(os.Stdout, "Ecommerce-Api 👉 ", log.LstdFlags)

	// server handler
	servermux := http.NewServeMux()
	servermux.Handle("/", handlers.NewRoot(logs))
	servermux.Handle("/cart", handlers.NewCart(logs))

	// server
	server := &http.Server{
		Addr:         ":8000",
		Handler:      servermux,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		logs.Printf("🚀 Listening to %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	signalchannel := make(chan os.Signal)
	signal.Notify(signalchannel, os.Interrupt)
	signal.Notify(signalchannel, os.Kill)
	sig := <-signalchannel
	logs.Println("🔻 Recived termination : shutting down api... ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	server.Shutdown(ctx)
}
