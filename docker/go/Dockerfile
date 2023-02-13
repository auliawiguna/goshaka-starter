FROM golang:1.19.0

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go mod download && go mod verify
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN GOOS=linux make watch

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080