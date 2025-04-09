package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type Request struct {
	url        string
	statuscode int
}

func main() {

	handler()
}

func handler() {
	start := time.Now()
	var url *string
	var requests *int
	var concurrency *int

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
	for ;counter < totalRequests; counter = counter + totalConcurrency{
		if  totalRequests < counter+totalConcurrency {
			break
		}
	}
	rest =  totalRequests - counter
	fmt.Printf("counter is %d\n", counter)
	fmt.Printf("rest is %d\n", rest)
	i := 0
	for i < counter {
		req, err := http.Get(*url)
		if err != nil {
			results <- Request{url: *url, statuscode: req.StatusCode}
		}
		results <- Request{url: *url, statuscode: req.StatusCode}
		i++
	}

	for rest > 0 {
		req, err := http.Get(*url)
		if err != nil {
			results <- Request{url: *url, statuscode: req.StatusCode}
		}
		results <- Request{url: *url, statuscode: req.StatusCode}
		rest --
	}

	sucesses := 0
	errors := 0

	for range totalRequests {
		result := <-results

		if result.statuscode == 200 {
			sucesses++
		}
		if result.statuscode != 200 {
			errors++
		}

	}

	end := time.Since(start)

	fmt.Println("URL:", *url)
	fmt.Println("Número total de requests:", *requests)
	fmt.Println("Número de chamadas simultâneas:", *concurrency)
	fmt.Println("Quantidade de requests com status Executados com sucesso:", sucesses)
	fmt.Println("Quantidade de requests com status Erro:", errors)
	fmt.Println("Tempo total gasto na execução:", end)
}
