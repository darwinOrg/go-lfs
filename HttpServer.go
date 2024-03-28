package main

import (
	"fmt"
	"github.com/rs/cors"
	fs "lfs/file-store"
	"lfs/setting"
	"log"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	router.POST("/upload", fs.Upload)
	router.GET("/download", fs.Download)

	h2s := &http2.Server{}

	host := fmt.Sprintf("0.0.0.0:%d", setting.GetAppInfo().Port)

	server := &http.Server{
		Addr:        host,
		Handler:     forCors(h2c.NewHandler(router, h2s)),
		IdleTimeout: time.Minute * 30,
	}

	log.Printf("Listening [0.0.0.0:8080]...\n")

	checkErr(server.ListenAndServe(), "while listening")
}

func forCors(h http.Handler) http.Handler {
	conf := setting.GetCorsConf()
	if conf == nil {
		return h
	}
	c := cors.New(cors.Options{
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: conf.Debug,
	})
	return c.Handler(h)
}

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	log.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}
