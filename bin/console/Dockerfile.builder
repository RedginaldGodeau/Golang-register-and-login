FROM golang:1.21.0-alpine AS build
WORKDIR /var/www/application

ENTRYPOINT ls | echo && go build -o ./app