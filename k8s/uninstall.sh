#!/bin/bash
kubectl delete deployment myapp-deployment -n allanfvc
kubectl delete service myapp -n allanfvc
kubectl delete namespace allanfvc