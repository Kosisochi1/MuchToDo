# MuchTodo Backend — Docker & Kubernetes

Containerization and orchestration of a Golang backend application using Docker, Docker Compose, and Kubernetes (Kind).
This project demonstrates how to move a backend service from a traditional server setup to a scalable, reproducible, and production-ready DevOps workflow.

## 📌 Project Overview

* Containerized a Golang backend application

* Orchestrated MongoDB, Redis, and the API using Docker Compose

* Built and pushed images to Docker Hub

* Deployed the application to a local Kubernetes cluster using Kind

* Exposed the service using NGINX Ingress

* Applied best practices for security, networking, and configuration management

# 🧰 Tech Stack

* Golang

* Docker

* Docker Compose

* Kubernetes (Kind)

* MongoDB

* Redis

* NGINX Ingress Controller
# 🛠 Prerequisites

Ensure the following are installed:

* Docker Desktop

* kubectl

* Kind (Kubernetes in Docker)

* Docker Hub account
# 📥 Clone Repository

 ```
 git clone https://github.com/Innocent9712/much-to-do.git
cd container-assessment
```

# ⚙ Setup Environment 
Copy the .env.example to .env and fill your values 

```
copy .env.example .env
```

# 🐳 Docker Setup
Build and Run with Docker Compose

```
docker-compose up --build
```
This will start:

* Backend API

* MongoDB

* Redis

* Redis Commander

Backend runs on:

```
http://localhost:8080
```

# 📦 Build & Push Docker Image
Login to Docker Hub:
```
docker login
```
Tag the image:

```
docker tag container-assessment-backend:latest kosisochi1/go-backend-api:latest
```

Push to Docker Hub:
```
docker push kosisochi1/go-backend-api:latest
```
# ☸️ Kubernetes Deployment (Kind)
Create Kind Cluster with port mapping as contain in the config file.
```
kind create cluster --name muchtodo-app --config kind-config.yaml
```
Install NGINX Ingress Controller
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```
# 📂 Kubernetes Manifests

Resources are organized as follows:
```
kubernates/
├── namespace.yaml
├── ingress.yaml
├── mongodb/
│   ├── mongodb-configmap.yaml
│   ├── mongodb-secret.yaml
│   ├── mongodb-deployment.yaml
│   ├── mongodb-service.yaml
│   └── mongodb-pvc.yaml
└── backend/
    ├── backend-configmap.yaml
    ├── backend-secret.yaml
    ├── backend-deployment.yaml
    └── backend-service.yaml
```

# 🚀 Deploy to Kubernetes
```
kubectl apply -f namespace.yaml
kubectl apply -f ingress.yaml
kubectl apply -f kubernates/mongodb
kubectl apply -f kubernates/backend
```
Verify:
```
kubectl get all -n muchtodo-app
```
# 🌐 Access the Application

Update your hosts file:
### Windows
```
C:\Windows\System32\drivers\etc\hosts
```
Add:
```
127.0.0.1 app-ingress.com

```
Open in browser:

```
http://app-ingress.com
```

# 🔄 Automate the process with Scripts
Run the following script in the terminal from the container-assessment root directory.

### Build image from the Dockerfile 
```
docker-build.sh
```

### Stert the service with docker-compose
```
docker-run.sh
```
### Deploy Kubernetes 
```
k8s-deploy.sh
```
### Deleting the kubernetes resources
```
k8s-cleanuo.sh
```



# 🔐 Key DevOps Practices Applied

* Multi-stage Docker builds for smaller images

* Non-root container execution

* Environment configuration via Kubernetes ConfigMaps & Secrets

* Health checks and resource limits

* Service-to-service communication using Kubernetes DNS

* Persistent storage for MongoDB
# ✅ Lessons Learned

* Avoid baking .env files into Docker images

* Kubernetes Secrets provide safer configuration management

* Debugging with kubectl is a critical Kubernetes skill

* Infrastructure as code reduces setup errors significantly

#   👤 Author

### Ezeoyiri Emmanuel Kosisochukwu

DevOps | Docker | Kubernetes | Golang

 

