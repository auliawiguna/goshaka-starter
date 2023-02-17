#!/bin/sh

docker-compose -f docker-compose.yml --env-file .docker-env down 
docker-compose -f docker-compose.yml --env-file .docker-env up -d 

# To rebuild
# docker-compose -f docker-compose.yml --env-file .docker-env up -d --build

# To force rebuild
# docker-compose -f docker-compose.yml --env-file .docker-env up -d --build --force-recreate 