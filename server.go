package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := flag.String("host", "localhost", "Host Address")
	port := flag.String("port", "8080", "Port")
	dir := flag.String("static", "./", "Path to static files")

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	log.Printf("%v:%v - Serve %v\n", *host, *port, *dir)
	http.ListenAndServe(fmt.Sprintf("%v:%v", *host, *port), nil)
}
