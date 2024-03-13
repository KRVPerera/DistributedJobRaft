#!/usr/bin/env bash

docker build . --build-arg CONFIG_FILE_PATH=config.xml -t djraft-node1:latest
docker build . --build-arg CONFIG_FILE_PATH=config2.xml -t djraft-node2:latest
docker build . --build-arg CONFIG_FILE_PATH=config3.xml -t djraft-node3:latest