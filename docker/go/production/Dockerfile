############################
# STEP 1 build executable binary
############################

# Digest SHA sha256:a00a03c72ecc1a01e4ca7b7cfb62459063912ff45ccb13011f9fb061e25916eb
# see https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
# FROM golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa as build
FROM golang:1.19.5-alpine3.17 AS build

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add make && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
# Create appuser
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"


WORKDIR $GOPATH/src/goshaka

COPY . .

RUN go get -d -v
# RUN go mod download && go mod verify

# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/goshaka

############################
# STEP 2 build a small image
############################
# FROM scratch
FROM alpine:latest

# Install supervisord and other required packages
RUN apk add --no-cache supervisor curl

# Copy the supervisord.conf file to the container
COPY ./docker/go/production/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

# Import the user and group files from the build.
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
# Copy our static executable.
COPY --from=build /go/bin/goshaka /go/bin/goshaka

COPY --from=build /go/src/goshaka/.env /go/bin/.env

# Create a directory for supervisord log files and set permissions
RUN mkdir -p /var/log/supervisor

# Use an unprivileged user.
# USER appuser:appuser
USER root

VOLUME /data

EXPOSE 3000

CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/conf.d/supervisord.conf"]

# Without supervisord
# ENTRYPOINT ["/go/bin/goshaka"]
