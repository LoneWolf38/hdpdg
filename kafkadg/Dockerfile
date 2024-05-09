FROM golang:1.22.3-alpine3.18
RUN  mkdir -p /hdpdg/kafka_go

ENV GOPATH=/hdpdg/
WORKDIR /hdpdg/kafka_go
COPY . /hdpdg/kafka_go

RUN go mod tidy && go mod vendor && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -ldflags='-w -s' -o /hdpdg/bin/kafka-dg-server

ENTRYPOINT [ "/hdpdg/kafka_go/bin/kafka-dg-server" ]

