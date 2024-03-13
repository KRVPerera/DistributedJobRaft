#!/usr/bin/env bash

docker run --network host -p 8082:8080 djraft-node2:latest
docker run --network host -p 8083:8080 djraft-node3:latest
docker run --network host -p 8081:8080 djraft-node1:latest