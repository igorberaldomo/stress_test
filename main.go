package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

// type Request struct {
// 	url        string
// 	statuscode int
// }

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

	results := make(chan int, *requests)

	getStatusIntoChannel(url, requests, concurrency, results)

	// <-results
	sucesses := 0
	errors404 := 0
	errors500 :=0 

	for range *requests {
		result := <-results
		if result == 200 {
			sucesses++
		}
		if result == 404  {
			errors404++
		}
		if result == 500 {
			errors500++
		}
	}

	end := time.Since(start)

	fmt.Println("URL:", *url)
	fmt.Println("Número total de requests:", *requests)
	fmt.Println("Número de chamadas simultâneas:", *concurrency)
	fmt.Println("Quantidade de requests com status Executados com sucesso:", sucesses)
	fmt.Println("Quantidade de requests com status Erro 404:", errors404)
	fmt.Println("Quantidade de requests com status Erro 500:", errors500)
	fmt.Println("Tempo total gasto na execução:", end)
}

func getStatusIntoChannel(url *string, requests *int, concurrency *int, results chan int) {

	counter := *requests / *concurrency
	rest := *requests - (*concurrency * counter)

	for range counter {
		for range *concurrency {
			go get(url, results)
		}
		// time.Sleep(time.Millisecond * 500)
	}
	if rest > 0 {
		for range rest {
			go get(url, results)
		}
	}
}

func get(url *string, results chan int) {
	req, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
		if req == nil {
			println("Nenhuma resposta recebida do servidor. Verifique a URL.")
			panic(err)
		}
		results <- req.StatusCode
		return
	}
	results <- req.StatusCode
}
