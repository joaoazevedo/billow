package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: 8080\n", *directory)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
