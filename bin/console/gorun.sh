#!/bin/bash

if [ ! -f "app" ]; then
  ./bin/console/gobuilder.sh
  exit 1
fi

docker compose up -d