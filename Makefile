all: docker_run

docker_run: docker_build_1 docker_build_2 docker_build_3
    docker run -p 8082:8080 djraft-node2:latest
    docker run -p 8083:8080 djraft-node3:latest
    docker run -it --rm -p 8080:8080 djraft-node1:latest

docker_build_1:
    docker build . --build-arg CONFIG_FILE_PATH=config.xml -t djraft-node1:latest

docker_build_2:
    docker build . --build-arg CONFIG_FILE_PATH=config2.xml -t djraft-node3:latest

docker_build_3:
    docker build . --build-arg CONFIG_FILE_PATH=config3.xml -t djraft-node3:latest

.PHONY: docker_run docker_build_1 docker_build_2 docker_build_3
