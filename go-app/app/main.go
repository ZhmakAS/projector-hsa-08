package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/tevjef/go-runtime-metrics/expvar"

	pkg_http "projector-test-app/pkg/http"
)

func main() {
	var cfg Env
	if err := cfg.Parse(); err != nil {
		panic(err)
	}

	faker := gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		osCall := <-c
		log.Printf("Stop system call:%+v", osCall)
		cancel()
	}()

	fs := http.FileServer(http.Dir("/app/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	server := pkg_http.NewServer(cfg.Addr, logRequest(http.DefaultServeMux))

	log.Printf("Starting up the API server...")
	if err := server.ListenAndServe(ctx); err != nil {
		log.Printf("Failed to serve api %s", err.Error())
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
