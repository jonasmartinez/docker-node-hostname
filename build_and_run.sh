#!/usr/bin/env bash

docker build -t jonasmartinez/docker-node-hostname .

docker rm -f docker-node-hostname

docker run -d \
       --name docker-node-hostname \
       -p 5000:5000 \
       jonasmartinez/docker-node-hostname

docker image prune -f
