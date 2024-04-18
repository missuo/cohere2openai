FROM golang:1.22 AS builder
WORKDIR /go/src/github.com/missuo/cohere2openai
COPY main.go ./
COPY go.mod ./
COPY go.sum ./
COPY types.go ./
COPY utils.go ./
RUN go get -d -v ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o cohere2openai .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/github.com/missuo/cohere2openai/cohere2openai /app/cohere2openai
CMD ["/app/cohere2openai"]