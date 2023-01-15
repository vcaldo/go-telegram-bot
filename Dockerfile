FROM golang:1.19-alpine

WORKDIR /app

COPY cmd/go-telegram-bot/go.mod ./
COPY cmd/go-telegram-bot/go.sum ./
RUN go mod download

COPY cmd/go-telegram-bot/*.go ./
RUN go build -o /bot

CMD [ "/bot" ]
