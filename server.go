package main

import (
	"CoursePlatform/Auth"

	"CoursePlatform/Base"
	"os"
	"net/http"
	"time"
	"os/signal"
	"context"
)


const srvAddr = "127.0.0.1:16001"

func main() {
	go Auth.StartServer()

	Base.SetupLoggerConfig()

	Base.Log.Notice("REST SERVER start\nPress Ctrl+C to shutdown")
	router := InitRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         srvAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			Base.Log.Info(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	var wait time.Duration
	wait = time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	Base.Log.Notice("REST shutting down")
	os.Exit(0)

}
