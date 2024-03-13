#!/usr/bin/env bash

docker build . --build-arg CONFIG_FILE_PATH=config.xml -t djraft-node1:latest
docker build . --build-arg CONFIG_FILE_PATH=config2.xml -t djraft-node2:latest
docker build . --build-arg CONFIG_FILE_PATH=config3.xml -t djraft-node3:latest

docker run -p 8082:8080 djraft-node2:latest
docker run -p 8083:8080 djraft-node3:latest
docker run -it --rm -p 8080:8080 djraft-node1:latest