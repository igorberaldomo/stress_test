package main

import (
	"flag"
	"fmt"
	"net/http"
)
type Request struct {
	url string
	statuscode int
}
func main() {
	var url *string
	var requests *int
	var concurrency *int
	handler(url, requests, concurrency)
}


func handler(url *string, requests *int, concurrency *int) {

	
	url = flag.String("url", "http://localhost:8080", "URL do serviço a ser testado.")
	if *url == "" {
		fmt.Println("Por favor, fornecer uma URL.")
		flag.PrintDefaults()
		return
	}
	requests = flag.Int("requests", 135, "Número total de requests.")
	concurrency = flag.Int("concurrency", 6, "Número de chamadas simultâneas.")

	flag.Parse()
	totalRequests := *requests
	totalConcurrency := *concurrency
	counter := 0
	rest := 0
	results := make(chan Request, totalRequests) 
	for counter <  totalRequests {
		counter = counter + totalConcurrency
	}
	rest =  counter - totalRequests

	for range counter{
		req , err := http.Get(*url)
		if err != nil {
			results <- Request{url: *url, statuscode: req.StatusCode}
		}
		results <- Request{url: *url, statuscode: req.StatusCode}
	}

	if rest > 0 {
		for range rest{
			req , err := http.Get(*url)
			if err != nil {
				results <- Request{url: *url, statuscode: req.StatusCode}
			}
			results <- Request{url: *url, statuscode: req.StatusCode}
		}
	}
	


	fmt.Println("URL:", *url)
	fmt.Println("Número total de requests:", *requests)
	fmt.Println("Número de chamadas simultâneas:", *concurrency)
}