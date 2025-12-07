#!/bin/bash

set -e

IMAGE="namduong0606/golang-portal:latest"

echo "➡️ Build Docker image..."
docker build -t $IMAGE .

echo "➡️ Push image..."
docker push $IMAGE

echo "➡️ Restart K8s deployment..."
kubectl rollout restart deployment/golang-portal

echo "➡️ Checking rollout status..."
kubectl rollout status deployment/golang-portal

echo "✔️ Done!"
