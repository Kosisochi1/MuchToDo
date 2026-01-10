#!/usr/bin/env bash

set -e 

NAMESPACE=muchtodo-app

echo "🚮 Deleting Kurbernetes resources..."

kubectl delete namespace $NAMESPACE --ignore-not-found


echo "✅ Cleanup Completed"

clear