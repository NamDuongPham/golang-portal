#!/usr/bin/env bash
# run.sh - build image, push to registry, and deploy to Kubernetes
# Compatible with Windows (Git Bash, WSL) and Unix shells
# Usage: ./run.sh [--tag TAG] [--minikube] [--build-only] [--push-only] [--deploy-only] [--namespace NS]

set -euo pipefail

# Defaults
IMAGE="${IMAGE:-namduong0606/golang-portal}"
TAG="${TAG:-latest}"
NAMESPACE="${NAMESPACE:-default}"
MINIKUBE=false
BUILD_ONLY=false
PUSH_ONLY=false
DEPLOY_ONLY=false

# Parse arguments
while [[ $# -gt 0 ]]; do
  case "$1" in
    --tag) TAG="$2"; shift 2;;
    --minikube) MINIKUBE=true; shift;;
    --build-only) BUILD_ONLY=true; shift;;
    --push-only) PUSH_ONLY=true; shift;;
    --deploy-only) DEPLOY_ONLY=true; shift;;
    --namespace) NAMESPACE="$2"; shift 2;;
    --help|-h) grep "^# " "$0" | head -15; exit 0;;
    *) echo "Unknown arg: $1"; exit 1;;
  esac
done

FULL_IMAGE="$IMAGE:$TAG"

echo "=== Deploy Config ==="
echo "IMAGE=$FULL_IMAGE"
echo "NAMESPACE=$NAMESPACE"
echo "MINIKUBE=$MINIKUBE"
echo ""

# Build stage
if [ "$DEPLOY_ONLY" = false ] && [ "$PUSH_ONLY" = false ]; then
  if [ "$MINIKUBE" = true ]; then
    if command -v minikube >/dev/null 2>&1; then
      echo "Building image inside Minikube: $FULL_IMAGE"
      minikube image build -t "$FULL_IMAGE" .
    else
      echo "minikube not found, using docker build"
      docker build -t "$FULL_IMAGE" .
    fi
  else
    echo "Building docker image: $FULL_IMAGE"
    docker build -t "$FULL_IMAGE" .
  fi
fi

if [ "$BUILD_ONLY" = true ]; then
  echo "Build-only mode. Done."
  exit 0
fi

# Push stage
if [ "$MINIKUBE" = false ] && [ "$DEPLOY_ONLY" = false ]; then
  echo "Pushing image: $FULL_IMAGE"
  docker push "$FULL_IMAGE"
fi

if [ "$PUSH_ONLY" = true ]; then
  echo "Push-only mode. Done."
  exit 0
fi

# Deploy stage
if [ "$BUILD_ONLY" = false ]; then
  echo "Deploying to namespace: $NAMESPACE"
  if kubectl set image deployment/golang-portal golang-portal="$FULL_IMAGE" -n "$NAMESPACE"; then
    echo "Triggered set-image"
  else
    echo "set-image failed, trying rollout restart"
    kubectl rollout restart deployment/golang-portal -n "$NAMESPACE"
  fi
  kubectl rollout status deployment/golang-portal -n "$NAMESPACE"
  kubectl get pods -l app=golang-portal -n "$NAMESPACE" -o wide
fi

echo "Done."
