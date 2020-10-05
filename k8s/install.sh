#!/bin/bash
kubectl create -f namespace.yml
kubectl create -f deployment.yml
kubectl create -f service.yml
if command -v minikube &> /dev/null
then
  minikube service list
fi