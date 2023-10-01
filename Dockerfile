FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD *.go ./

RUN go build


FROM alpine:3.18

WORKDIR /

COPY --from=build /app/rowenta-robot-vacuum-exporter .

USER nobody

ENTRYPOINT [ "/rowenta-robot-vacuum-exporter" ]