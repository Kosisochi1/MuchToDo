#!/usr/bin/env bash

# Script to build and push to registry

# Stop script on error
set -e

APP_NAME=kosisochi1/go-backend-api
IMAGE_TAG=latest

echo "Building Docker Image...."

docker build -t ${APP_NAME}:${IMAGE_TAG} -f Dockerfile .
echo "🐋 Image built: ${APP_NAME}:${IMAGE_TAG}"


# Ensure you are Logged in to Dockerhub.

echo "Pushing 🐋 Image to Dockerhub"

docker push kosisochi1/go-backend-api:latest

