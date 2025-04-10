FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go mod tidy
RUN CGO_ENABLED=0 CGOOS=linux GOARCH=amd64 go build -o main main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/main .
ENTRYPOINT ["./main"]


# docker build -t stresstester .

# docker run --name stresstester stresstester --url=http://172.17.0.2:8080 --requests=105 --concurrency=10