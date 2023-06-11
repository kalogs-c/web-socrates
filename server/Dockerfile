# Change this variable to the name of the app
ARG APP=socrates-server

ARG GO_VERSION=1.20.1

######################################

FROM golang:$GO_VERSION-alpine as builder

WORKDIR /$APP

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o $APP .

# Install ca-certificates, remove the next line if not needed
# Some databases (like Postgres) require ca-certificates
# And scratch image does not have included
RUN apk add -U --no-cache ca-certificates

######################################

FROM scratch

COPY --from=builder /$APP/$APP .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT [ "./$APP" ]
