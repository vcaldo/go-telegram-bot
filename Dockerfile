FROM golang:1.19-alpine

WORKDIR /app

COPY bot/go.mod ./
COPY bot/go.sum ./
RUN go mod download

COPY bot/*.go ./
RUN go build -o /bot

CMD [ "/bot" ]
