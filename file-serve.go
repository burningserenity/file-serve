// Package file-serve implements a simple http server which serves content from the current working directory
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type FileSystem interface {
	Open(name string) (http.File, error)
}

func main() {
	r := http.NewServeMux()
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	port := ":" + "8000"

	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	var fs FileSystem = http.Dir(cwd)

	serve := http.FileServer(fs)
	r.Handle("/", handlers.LoggingHandler(os.Stdout, serve))

	log.Fatal(http.ListenAndServe(port, handlers.CompressHandler(r)))
}
