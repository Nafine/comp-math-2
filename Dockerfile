FROM golang:1.26rc1-alpine3.23  AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/comp-math-2/comp-math-2.go

FROM alpine:3.23

COPY --from=builder /app/server /app/server

EXPOSE 8080
ENTRYPOINT ["/app/server"]