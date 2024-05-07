FROM golang:latest

WORKDIR /usr/src/app
ENV PORT 3000
ENV HOST 0.0.0.0

COPY . .

RUN go build -o bin .

CMD go run main.go

ENTRYPOINT [ "/usr/src/app/bin" ]
