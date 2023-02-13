############################
# STEP 1 build executable binary
############################

# Digest SHA sha256:a00a03c72ecc1a01e4ca7b7cfb62459063912ff45ccb13011f9fb061e25916eb
# see https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
FROM golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa as builder
# FROM golang:1.19.5-alpine3.17 AS build

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
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


WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify

RUN go install github.com/swaggo/swag/cmd/swag@latest

# Development Mode
RUN GOOS=linux make watch