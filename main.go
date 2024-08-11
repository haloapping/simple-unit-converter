package main

import (
	"fmt"
	"net/http"
)

const (
	host = "localhost"
	port = "8000"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", IndexPageHandler)
	mux.HandleFunc("GET /length.html", LengthPageHandler)
	mux.HandleFunc("GET /weight.html", WeightPageHandler)
	mux.HandleFunc("GET /temperature.html", TemperaturePageHandler)

	mux.HandleFunc("POST /length.html", ConvertHandler)
	mux.HandleFunc("POST /weight.html", ConvertHandler)
	mux.HandleFunc("POST /temperature.html", ConvertHandler)

	fmt.Printf("Server started at %v:%v\n", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), mux); err != nil {
		panic(err)
	}
}
