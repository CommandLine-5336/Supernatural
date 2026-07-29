FROM golang:1.26.5-alpine AS build
WORKDIR /mail_service
COPY . .
# hadolint ignore=DL3008
RUN apk add --no-cache tzdata && go build -o main

# Runtime stage
FROM alpine:3.23.5
WORKDIR /auth
COPY --from=build /mail_service/main ./
COPY --from=build /mail_service/templates ./templates
EXPOSE 8080
CMD ["./main"]
