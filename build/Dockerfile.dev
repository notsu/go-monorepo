FROM golang:1.11.1-alpine3.8

# install git, curl and realize
RUN apk add --no-cache git curl && \
  go get -u github.com/tockins/realize && \
  go get -u github.com/derekparker/delve/cmd/dlv
