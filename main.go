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
	mux.HandleFunc("GET /index.html", IndexPageHandler)
	mux.HandleFunc("GET /pages/length.html", LengthPageHandler)
	mux.HandleFunc("GET /pages/weight.html", WeightPageHandler)
	mux.HandleFunc("GET /pages/temperature.html", TemperaturePageHandler)

	mux.HandleFunc("POST /pages/length.html", ConvertLengthHandler)
	mux.HandleFunc("POST /pages/weight.html", ConvertWeightHandler)
	mux.HandleFunc("POST /pages/temperature.html", ConvertTemperatureHandler)

	fmt.Printf("Server started at %v:%v\n", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), mux); err != nil {
		panic(err)
	}
}
