FROM golang:1.16-alpine

WORKDIR /app

ENV DOK8S_ROLE="CONSUMER"

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o do-k8s-kafka .

ENTRYPOINT [ "/app/do-k8s-kafka" ]