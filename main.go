package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"bufio"
	"os"
	"strconv"

)

func main() {
	var url *string
	var requests *int
	var concurrency *int

	handler(url, requests, concurrency	)
}

func handler(url *string, requests *int, concurrency *int) {
	start := time.Now()

	var URLCMD string
	var RequestsCMD string
	var ConcurrencyCMD string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Por favor, fornecer a URL do serviço a ser testado.")
	URLCMD, _ = reader.ReadString('\n')
	if URLCMD == "" {
		URLCMD = "none"
	}
	fmt.Println("Por favor, fornecer o número total de requests.")
	RequestsCMD, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	if RequestsCMD == "" {
		RequestsCMD = "0"
	}
	requestscmd, _ := strconv.Atoi(RequestsCMD)

	fmt.Println("Por favor, fornecer o número de chamadas simultâneas.")
	ConcurrencyCMD, _ = reader.ReadString('\n')
	if ConcurrencyCMD == "" {
		ConcurrencyCMD = "0"
	}
	concurrencycmd, _ := strconv.Atoi(ConcurrencyCMD)

	url = flag.String("url", "http://localhost:8080", "URL do serviço a ser testado.")
	requests = flag.Int("requests", 135, "Número total de requests.")	
	concurrency = flag.Int("concurrency", 6, "Número de chamadas simultâneas.")
	if *url == "" {
		fmt.Println("Por favor, fornecer uma URL.")
		flag.PrintDefaults()
		return
	}


	flag.Parse()

	results := make(chan int, *requests)

	getStatusIntoChannel(url ,URLCMD, requests,requestscmd, concurrency,concurrencycmd, results)


	sucesses := 0
	errors := 0

	for range *requests {
		result := <-results
		if result == 200 {
			sucesses++
		}
		if result != 200 {
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

func getStatusIntoChannel(url *string, urlcmd string, requests *int, requestscmd int, concurrency *int, concurrencycmd int, results chan int) {
	fmt.Printf("teste")

	counter := *requests / *concurrency
	rest := *requests - (*concurrency * counter)

	for range counter {
		for range *concurrency {
			go get (url, results)
		}
	}
	if rest > 0 {
		for range rest {
			go get (url, results)
		}
	}
}

func get (url *string, results chan int) {
	req, err := http.Get(*url)
	if err != nil {
		results <- req.StatusCode
		return
	}
	results <- req.StatusCode
}