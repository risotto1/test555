#!/bin/sh
kubectl config set-credentials jenkins --token="${JENKINS_SA_TOKEN}"
kubectl config set-cluster minikube --server="https://192.168.99.100:8443" --certificate-authority="${APISERVER_CA}"
kubectl config set-context jenkins --cluster minikube --user=jenkins --namespace=kube-system
kubectl config use-context jenkins
chmod 755 ~/.kube/config
