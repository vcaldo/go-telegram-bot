FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY qbitorrent qbitorrent/
RUN go build -o /bot

CMD [ "/bot" ]
