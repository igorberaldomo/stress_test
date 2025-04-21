# Stress Tester

Objetivo fazer um projeto que vai realizar um número de requisições em um endereço usando un numero determinado de chamadas simultâneas. Depois disso ele vai imprimir na tela um relatório

```
    Apresentar um relatório ao final dos testes contendo:
    Tempo total gasto na execução
    Quantidade total de requests realizados.
    Quantidade de requests com status HTTP 200.
    Distribuição de outros códigos de status HTTP (como 404, 500, etc.). // representados como StatusNotFound
```

Para fazer o build do `server` e do `stress_tester` execute no bash no raiz do projeto

```bash
    make build_all
```

Para executar o docker container do `server` execute no bash no raiz do projeto


```bash
    make run_server
```

Para executar o docker container do `stress_tester` execute no bash no raiz do projeto

```bash
    make run_stress_tester
```

Como deve ser a saida do relatório:
```
    URL: http://localhost:8080
    Número total de requests: 135
    Número de chamadas simultâneas: 6
    Quantidade de requests com status Executados com sucesso: 127
    Quantidade de requests com status Erro: 8
    Tempo total gasto na execução: 20.607732ms
```
