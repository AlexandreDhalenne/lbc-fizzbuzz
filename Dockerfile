FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /lbc-fizzbuzz cmd/main.go
EXPOSE 8080
ENV GIN_MODE=release
CMD ["/lbc-fizzbuzz"]