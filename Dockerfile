# step 1 - build executable binary
FROM golang:1.13-alpine as builder

WORKDIR /app

# fetch base packages & create a default user & copy project files
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' appuser

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .

# expose echo http service
EXPOSE 8080

# fetch go dependencies
RUN go mod download

COPY . .

RUN go build -o gorsvp

# step 2 - build compact image
FROM scratch

# import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# copy our static executable
COPY --from=builder /app/gorsvp /go/bin/gorsvp

# use an unprivileged user
USER appuser

# expose echo http service
EXPOSE 8080

# run our app binary.
ENTRYPOINT ["/go/bin/gorsvp"]
