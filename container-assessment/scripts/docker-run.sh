#!/usr/bin/env bash

set -e 


echo "⚡ Starting service with Docker Compose..."

docker-compose up -d --build 

echo " Services are running"

docker-compose ps


clear