package Auth

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/op/go-logging"

	"net/http"
)

var log = logging.MustGetLogger("Auth Logger")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

const srvAddr = "127.0.0.1:16000"

func StartSer+ver() {
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)

	log.Notice("AUTH SERVER start\nPress Ctrl+C to shutdown")
	router := InitRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         srvAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info(err)
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

	log.Notice("auth shutting down")
	os.Exit(0)

}
