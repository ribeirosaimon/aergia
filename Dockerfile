FROM golang:latest as builder
WORKDIR /app
COPY . /app
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o aergia /app/src/main.go

FROM alpine

LABEL maintainer = "aergiaApp"

WORKDIR /app
COPY --from=builder /app/aergia /app
ENTRYPOINT ./aergia