FROM golang:alpine AS build
WORKDIR /go/src/ovpnstatusd
COPY . .
RUN go build

FROM alpine:latest
COPY --from=build go/src/ovpnstatusd/ovpnstatusd /usr/local/bin
ENTRYPOINT ["ovpnstatusd"]
