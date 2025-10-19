#! /bin/bash

## sudo chown -R 1000:1000 logs/kafka-*

EXTERNAL_UP=$(curl ifconfig.me)

export DOCKER_HOST_IP=$EXTERNAL_UP

docker-compose -f docker-compose.kafka.yml up --build -d