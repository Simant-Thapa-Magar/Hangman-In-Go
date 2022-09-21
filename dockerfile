FROM golang:latest
WORKDIR /app
COPY go.mod ./
COPY *.* ./
COPY /states ./states
RUN go build -o /hangman-go
ENTRYPOINT ["/hangman-go"]