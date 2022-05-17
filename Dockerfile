# docker build -t go-simple-http .
# docker run -dit --name go-simple-http -p 8900:8900 go-simple-http

FROM golang:1.17 As builder
WORKDIR /app

COPY . .
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o go-simple-http

CMD ["/app/go-simple-http"]