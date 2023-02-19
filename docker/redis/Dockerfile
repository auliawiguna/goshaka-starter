FROM redis:alpine

# For the security reason, it is recommended to use redis.conf, kindly update the password within the redis.conf
# Create directory /usr/local/etc/redis and copy the redis.conf into that directory
# RUN mkdir -p /usr/local/etc/redis
# COPY redis.conf /usr/local/etc/redis/redis.conf

VOLUME /data

EXPOSE 6379

#CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
CMD ["redis-server"]










