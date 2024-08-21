#!/usr/bin/env bash

kubectl apply -f deployment/deployment.yaml -n default
kubectl apply -f deployment/service.yaml -n default