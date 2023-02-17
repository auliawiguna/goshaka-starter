#!/bin/sh

docker-compose -f docker-compose.yml --env-file .docker-env down 
docker-compose -f docker-compose.yml --env-file .docker-env up -d 

# Restart supervisord
# docker exec -it goshaka_be supervisorctl restart goshaka-worker:

# Remove cache
# docker builder prune

# To rebuild
# docker-compose -f docker-compose.yml --env-file .docker-env up -d --build

# To force rebuild
# docker-compose -f docker-compose.yml --env-file .docker-env up -d --build --force-recreate 

# To rebuild a service
# docker-compose -f docker-compose.yml --env-file .docker-env up -d --no-deps --build <service>