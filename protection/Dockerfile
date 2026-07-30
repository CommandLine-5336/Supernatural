# Build stage
FROM golang:1.26.5-alpine AS build
WORKDIR /protection
COPY go.mod go.sum main.go ./
RUN go build -o main

# Runtime stage
FROM alpine:3.23.5
WORKDIR /protection
COPY --from=build /protection/main ./
CMD ["sh", "-c", "while true; do /protection/main; sleep 60; done"]
