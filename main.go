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

	results := make(chan Request, *requests)

	go getURLIntoChannel(url, requests, concurrency, results)

	<-results
	sucesses := 0
	errors := 0

	for range *requests {
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

func getURLIntoChannel(url *string, requests *int, concurrency *int, results chan <-Request) {

	counter := 0
	rest := 0

	for ;counter < *requests; counter = counter + *concurrency{
		if  *requests < counter+*concurrency {
			break
		}
	}
	rest =  *requests - counter
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

}