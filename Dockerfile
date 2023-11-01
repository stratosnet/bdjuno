FROM golang:1.19-alpine AS builder
RUN apk update
RUN apk add --no-cache make git build-base gmp-dev flex bison

ARG pbc_lib_ver=0.5.14
RUN wget https://crypto.stanford.edu/pbc/files/pbc-${pbc_lib_ver}.tar.gz && \
  tar xzvf pbc-${pbc_lib_ver}.tar.gz && \
  cd pbc-${pbc_lib_ver} && \
  ./configure && \
  make && \
  make install && \
  ldconfig / && \
  cd $BUILD && \
  rm -rf pbc-${pbc_lib_ver}*

WORKDIR /go/src/github.com/forbole/bdjuno
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
COPY --from=builder /usr/local/lib/libpbc.so.1.0.0 /usr/local/lib/libpbc.so.1.0.0

RUN apk add --no-cache gmp-dev
RUN cd /usr/local/lib && { ln -s -f libpbc.so.1.0.0 libpbc.so.1 || { rm -f libpbc.so.1 && ln -s libpbc.so.1.0.0 libpbc.so.1; }; } \
  && cd /usr/local/lib && { ln -s -f libpbc.so.1.0.0 libpbc.so || { rm -f libpbc.so && ln -s libpbc.so.1.0.0 libpbc.so; }; }

CMD [ "bdjuno" ]
