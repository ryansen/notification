# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.11-alpine3.7 as builder

# intall tools
RUN apk add --no-cache git

# install /usr/bin/nsenter
RUN apk add --no-cache util-linux

WORKDIR /go/src/openpitrix.io/notification/
COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN  go build -v  -a -installsuffix cgo -ldflags '-w'  -o  cmd/nf_server_main  cmd/server/server_main.go



FROM alpine
COPY --from=builder /go/src/openpitrix.io/notification/cmd/nf_server_main /usr/local/bin/

EXPOSE 50051
CMD ["/usr/local/bin/nf_server_main"]


