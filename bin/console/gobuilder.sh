#!/bin/bash

if [ ! -f "go.mod" ]; then
  echo "go.mod does'nt exist"
  exit 1
fi

echo -e "Build Dockerfile :\n##################################"
docker build -f ./bin/console/Dockerfile.builder -t go_builder .

echo -e "Build Go Project :\n##################################"
docker run -t --rm -v ./:/var/www/application/  go_builder

clear
echo "Go Project Builded"