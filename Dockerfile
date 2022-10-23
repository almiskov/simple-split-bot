##
##  build
##

FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /bot

##
##  deploy
##

FROM gcr.io/distroless/base

WORKDIR /

COPY --from=builder /bot /bot
COPY ./config.json /config.json

ENTRYPOINT [ "/bot" ]