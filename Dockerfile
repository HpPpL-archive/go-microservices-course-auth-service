FROM golang:1.20.3-alpine AS builder

COPY . /github.com/HpPpL/go-microservices-course-auth-service/grpc/source/
COPY go.mod /github.com/HpPpL/go-microservices-course-auth-service/grpc/source/
COPY go.sum /github.com/HpPpL/go-microservices-course-auth-service/grpc/source/
WORKDIR /github.com/HpPpL/go-microservices-course-auth-service/grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_auth_server ./cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/HpPpL/go-microservices-course-auth-service/grpc/source/bin/crud_auth_server .

CMD ["./crud_auth_server"]
