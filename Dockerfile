FROM golang:1.18

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./entrypoint.sh /app/entrypoint.sh
COPY . .
RUN mkdir -p /app/input
COPY ./input/logfile.log /app/input/logfile.log
RUN cd cmd && go build -o main . && chmod +x main
EXPOSE 8080

CMD ["./cmd/main"]
