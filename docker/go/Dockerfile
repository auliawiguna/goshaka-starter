# Use Alpine as the base image
FROM alpine:3.13

# Set the environment variable for Go version
ENV GO_VERSION=1.19.0

# Install required dependencies and set up Go
RUN apk add --update \
    curl \
    supervisor \
    nano \
    git \
    build-base \
  && curl -fsSL https://golang.org/dl/go$GO_VERSION.src.tar.gz -o go.tar.gz \
  && tar -xzf go.tar.gz \
  && rm go.tar.gz \
  && cd go/src \
  && ./make.bash \
  && cd / \
  && rm -rf go \
  && apk del curl \
  git \
  build-base

# Set the environment variables for Go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify

RUN go install github.com/swaggo/swag/cmd/swag@latest

# Development Mode
RUN GOOS=linux make watch

# Production Mode
# RUN GOOS=linux make build
# FROM alpine:3.10
# RUN apk --no-cache add ca-certificates
# WORKDIR /usr/bin
# RUN mkdir -p /go/src/bin/goshaka
# COPY /go/src/server /go/bin/goshaka
# COPY /go/src/.env /go/bin/goshaka
# EXPOSE 3000
# ENTRYPOINT /go/bin/goshaka/server