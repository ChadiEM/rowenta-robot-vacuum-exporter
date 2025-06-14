FROM golang:1.24-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD *.go ./

RUN go build


FROM alpine:3.22

WORKDIR /

COPY --from=build /app/rowenta-robot-vacuum-exporter .

USER nobody

ENTRYPOINT [ "/rowenta-robot-vacuum-exporter" ]