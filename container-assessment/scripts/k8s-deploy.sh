#!/usr/bin/env bash

set -e 

# Deploying Kubernetes 
NAMESPACE=muchtodo-app

echo " Creating namespace if does not exist..."


kubectl apply -f kubernates/namespace.yaml 


# Deploying MongoDB

echo "✅ Appying mongodb manifest"

kubectl apply -f kubernates/mongodb/mongodb-configmap.yaml
kubectl apply -f kubernates/mongodb/mongodb-secret.yaml
kubectl apply -f kubernates/mongodb/mongodb-pvc.yaml
kubectl apply -f kubernates/mongodb/mongodb-deployment.yaml
kubectl apply -f kubernates/mongodb/mongodb-service.yaml


# Wait for database to be Ready


echo "Waiting for MongoDB deployment..."

kubectl wait \
  --for=condition=available deployment/mongodb-deployment \
  -n muchtodo-app \
  --timeout=120s

echo "MongoDB is ready ✅"


# Deploying Backend 

echo "Apply ✅ Backend manifest"

kubectl apply -f kubernates/backend/backend-configmap.yaml
kubectl apply -f kubernates/backend/backend-secret.yaml
kubectl apply -f kubernates/backend/backend-deployment.yaml
kubectl apply -f kubernates/backend/backend-service.yaml


#Wait for backend deployment to be ready

echo "Waiting for backend deployment..."

kubectl wait \
  --for=condition=available deployment/backend \
  -n muchtodo-app \
  --timeout=120s

echo "Waiting for backend pods..."


kubectl wait \
  --for=condition=ready pod \
  -l app=backend \
  -n muchtodo-app \
  --timeout=120s

echo "Backend is fully ready ✅"
# Applying Ingress


# Check if Ingress Controller is installed
if kubectl get namespace ingress-nginx &> /dev/null; then
    echo "Ingress controller already installed"
else
    echo "Installing NGINX Ingress Controller..."
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
    echo "Waiting for Ingress Controller to be ready..."
    kubectl wait --namespace ingress-nginx \
      --for=condition=ready pod \
      --selector=app.kubernetes.io/component=controller \
      --timeout=120s
fi

echo "Apply ✅ Ingress"


kubectl apply -f kubernates/ingress.yaml


# Check Deployment status


echo "✅ Deployment completed "

kubectl get all -n $NAMESPACE
clear


# Access the API 

# Edit the etc/hosts file to Add the DNS mapped to the localhost
# echo '127.0.0.1 app-ingress.com' | sudo tee -a /etc/hosts 
curl http://app-ingress.com
curl http://app-ingress.com/health

