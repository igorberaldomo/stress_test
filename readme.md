Objetivo fazer um projeto que vai realizar um número de requisições em um endereço usando un numero determinado de chamadas simultâneas. Depois disso ele vai imprimir na tela um relatório

    Apresentar um relatório ao final dos testes contendo:
    Tempo total gasto na execução
    Quantidade total de requests realizados.
    Quantidade de requests com status HTTP 200.
    Distribuição de outros códigos de status HTTP (como 404, 500, etc.). // representados como StatusNotFound

para ativação do programa fora do docker:

    // ativa o servidor na porta http://localhost:8080 a cada requisição ele vai emular uma falha 5% das requisição
    1º go run server/server.go

    // ativa os requests
    2º go run main.go 

para ativação com o docker:
    na pasta server:

    // suba a imagem do servidor para o docker
    1º docker build -t server .

    // suba o container a partir da imagem docker
    2º docker run --name server -p 8080:8080 server

    no root:

    // suba a imagem do tester
    3º docker build -t stresstester .

    // você precisa colocar os seguintes parametros para que a imagem execute corretamente
        endereço do container do server:
            --url=http://172.17.0.2:8080
        Nº de request:
            --requests=100
        Nº de concorrencia:
            --concurrency=12
    exemplo:
    4º docker run --name stresstester stresstester --url=http://172.17.0.2:8080 --requests=105 --concurrency=10

Como deve ser a saida do relatório:
    URL: http://localhost:8080
    Número total de requests: 135
    Número de chamadas simultâneas: 6
    Quantidade de requests com status Executados com sucesso: 127
    Quantidade de requests com status Erro: 8
    Tempo total gasto na execução: 20.607732ms
