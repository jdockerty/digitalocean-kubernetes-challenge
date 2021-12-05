#!/bin/bash

# Assumes that you are already using the correct kube config file.
# Ref: https://strimzi.io/quickstarts/

kubectl create namespace kafka
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n kafka 