# syntax=docker/dockerfile:1

FROM golang:1.21.0-alpine AS build
WORKDIR /var/www/application

# GO BUILD PROJECT
RUN go install && \
    go build -o /var/www/application/bin/app /var/www/application/

# LAUNCH PROJECT
FROM alpine:latest AS final
COPY --from=build /var/www/application/bin/app /var/www/application/bin/app

ENTRYPOINT ["/var/www/application/bin/app"]