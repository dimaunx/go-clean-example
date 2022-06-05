FROM golang:1.17-bullseye as builder

# Create and change to the app directory.
WORKDIR /go/src

# Copy dependencies
COPY ./vendor ./vendor
COPY ./pkg ./pkg
COPY go.mod go.sum main.go ./

# Build
RUN env CGO_ENABLED=0 GOOS=linux go build -o main

# Copy artifacts to a clean image
FROM golang:1.17-alpine3.16

COPY --from=builder /go/src/main ./main
ENTRYPOINT [ "./main" ]