package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/ethicalapps/ucms/api"
	"github.com/ethicalapps/ucms/cms"
	"github.com/ethicalapps/ucms/cms/bolt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	fHost = flag.String("host", "127.0.0.1", "service host")
	fPort = flag.String("port", ":0", "service port")
	fDB   = flag.String("db", "ucms.db", "database file")
)

func main() {
	flag.Parse()

	store, err := bolt.New(*fDB)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	cms.Init(store)

	listener, err := net.Listen("tcp", *fPort)
	if err != nil {
		log.Fatal(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/healthz"))
	router.Mount("/api", api.Router())

	go func() {
		err = http.Serve(listener, router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("[uCMS]", "listening on", *fHost, port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	// TODO: deregister service

}
