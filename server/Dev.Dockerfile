ARG GO_VERSION=1.20.1

######################################

FROM golang:$GO_VERSION-alpine

WORKDIR /app

COPY . .

# Install build basics applications and compilers that go test uses
RUN apk add build-base

# Block the terminal
CMD [ "tail", "-f", "/dev/null" ]
