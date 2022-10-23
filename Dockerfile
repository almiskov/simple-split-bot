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
COPY --from=builder /app/assets /assets
COPY ./config.json /config.json

ENTRYPOINT [ "/bot" ]