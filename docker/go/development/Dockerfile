############################
# STEP 1 build executable binary
############################

# see https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

FROM golang:1.19.5-alpine3.17 AS build

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add make && apk add --no-cache git ca-certificates && update-ca-certificates
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

RUN wget https://github.com/swaggo/swag/releases/download/v1.8.10/swag_1.8.10_Linux_x86_64.tar.gz && \
    tar -xvzf swag_1.8.10_Linux_x86_64.tar.gz && \
    mv swag $GOPATH/bin

RUN wget https://github.com/cespare/reflex/releases/download/v0.3.1/reflex_linux_amd64.tar.gz && \
    tar -xvzf reflex_linux_amd64.tar.gz && \
    mv reflex_linux_amd64/reflex $GOPATH/bin

RUN wget https://github.com/golangci/golangci-lint/releases/download/v1.51.1/golangci-lint-1.51.1-linux-amd64.tar.gz && \
    tar -xvzf golangci-lint-1.51.1-linux-amd64.tar.gz && \
    mv golangci-lint-1.51.1-linux-amd64/golangci-lint $GOPATH/bin

# RUN go mod download && go mod verify

VOLUME /data

EXPOSE 3000

CMD ["make", "watch"]