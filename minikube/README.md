Minikube
========

This folder contains a set of experiments I'm developing to
study a bit more about Kubernetes

The objective is to create a few applications that will work together
and deploy the whole infrastructure using Kubernetes (minikube for simplicity),
I'll probably have:
 * Backend
 * Frontend
 * Some cache/no-SQL database (probably Redis)
 * Postgres
 
Things I will experiment with:
 * Multiple containers on a Pod
 * Multiple applications/Pods working together (backend + frontend)
 * Redis/Postgres running on Kubernetes
 * Jobs/Cron Jobs
 * Sending logs to Kibana/GrayLog
