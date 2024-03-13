#!/usr/bin/env bash

docker inspect -f '{{range $key, $value := .NetworkSettings.Networks}}{{$key}} {{end}}' [container]
docker inspect -f '{{.NetworkSettings.Networks.[network].IPAddress}}' [container]
docker network inspect -f '{{range .Containers}}{{.Name}} {{end}}' [network]
docker network connect [network] [container]