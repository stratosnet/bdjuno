FROM golang:1.16-alpine AS builder
RUN apk update
RUN apk add --no-cache make git build-base
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
CMD [ "bdjuno" ]
