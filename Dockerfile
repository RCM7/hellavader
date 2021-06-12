FROM golang:1.14-alpine

ENV WORKDIR /app

WORKDIR $WORKDIR

COPY go.mod $WORKDIR
COPY go.sum $WORKDIR

RUN go mod download

COPY . $WORKDIR

RUN go build -o hellavader /app/hellavader.go

EXPOSE 8081

CMD $WORKDIR/hellavader

