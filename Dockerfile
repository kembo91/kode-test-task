FROM golang:alpine as buider
USER root
WORKDIR /app

COPY ./client ./client
COPY ./server ./server

RUN apk add --no-cache --update nodejs yarn

RUN mkdir build
RUN cd ./client; yarn install; yarn build
RUN mv ./client/build ./build/build
RUN cd ./server; go mod download; GOOS=linux GOARCH=amd64 go build main.go; mv main ../build

FROM alpine:latest

USER root

COPY --from=buider /app/build /app/build
COPY --from=buider /app/server/config /app/build/config

WORKDIR /app/build

ENTRYPOINT [ "./main" ]

EXPOSE 8080