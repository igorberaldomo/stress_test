FROM golang:1.24
RUN go mod tidy
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 CGOOS=linux GOARCH=amd64 go build -o main main.go
