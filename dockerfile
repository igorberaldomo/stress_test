FROM golang:1.24 AS server
WORKDIR /app
COPY server/server.go ./
RUN go mod tidy
RUN CGO_ENABLED=1 CGOOS=linux GOARCH=amd64 go build -o server.go server.go

FROM scratch
WORKDIR /app
COPY --from=server /app/server.go .
ENTRYPOINT ["./server.go"]
EXPOSE 8080

FROM golang:1.24 AS build
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      build-essential \
      libsqlite3-dev
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . ./
RUN CGO_ENABLED=1 CGOOS=linux GOARCH=amd64 go build -o main.go main.go

FROM debian:bookworm
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      libsqlite3-0
COPY --from=build /app/main.go /usr/bin/main.go
ENTRYPOINT ["main.go"]
CMD ["--url=http://localhost:8080", "--requests=100", "--concurrency=10"]