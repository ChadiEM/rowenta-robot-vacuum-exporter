FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o rowenta-robot-vacuum-exporter -ldflags "-s -w" /app/cmd/rowenta-robot-vacuum-exporter

FROM alpine:3.22

WORKDIR /

COPY --from=build /app/rowenta-robot-vacuum-exporter .

USER nobody

ENTRYPOINT [ "/rowenta-robot-vacuum-exporter" ]