FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o app
EXPOSE 7000
ENTRYPOINT ["./app"]

