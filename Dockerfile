FROM ubuntu:20.04 AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends  make git build-essential ca-certificates libgmp-dev flex bison curl wget \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ARG pbc_lib_ver=0.5.14
RUN wget https://crypto.stanford.edu/pbc/files/pbc-${pbc_lib_ver}.tar.gz \
    && tar xzvf pbc-${pbc_lib_ver}.tar.gz \
    && cd pbc-${pbc_lib_ver} \
    && ./configure \
    && make \
    && make install \
    && ldconfig / \
    && cd .. && rm -rf pbc-${pbc_lib_ver}*

RUN wget https://go.dev/dl/go1.22.12.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.12.linux-amd64.tar.gz \
    && ln -s /usr/local/go/bin/go /usr/bin/go \
    && rm go1.22.12.linux-amd64.tar.gz

WORKDIR /go/src/github.com/forbole/callisto
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN make build

FROM ubuntu:20.04
WORKDIR /callisto
COPY --from=builder /go/src/github.com/forbole/callisto/build/callisto /usr/bin/callisto
COPY --from=builder /go/src/github.com/forbole/callisto/database/schema /usr/share/callisto/database/schema
COPY --from=builder /go/src/github.com/forbole/callisto/hasura /usr/share/callisto/hasura
COPY --from=builder /usr/local/lib/libpbc.so.1.0.0 /usr/local/lib/libpbc.so.1.0.0

RUN cd /usr/local/lib && { ln -s -f libpbc.so.1.0.0 libpbc.so.1 || { rm -f libpbc.so.1 && ln -s libpbc.so.1.0.0 libpbc.so.1; }; } \
    && cd /usr/local/lib && { ln -s -f libpbc.so.1.0.0 libpbc.so || { rm -f libpbc.so && ln -s libpbc.so.1.0.0 libpbc.so; }; } \
    && ldconfig

CMD [ "callisto" ]
