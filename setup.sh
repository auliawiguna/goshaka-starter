#!/bin/sh

docker-compose -f docker-compose.yml --env-file .docker-env down 
docker-compose -f docker-compose.yml --env-file .docker-env up -d 