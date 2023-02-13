FROM nginx:alpine

COPY nginx.conf /etc/nginx/
COPY default.conf /etc/nginx/conf.d/

RUN apk update \
    && apk upgrade \
    && apk --update add logrotate \
    && apk add --no-cache openssl \
    && apk add --no-cache bash

RUN apk add --no-cache curl

EXPOSE 80 81 443
















