#!/bin/bash

kubectl -n tekton-pipelines create secret generic github-credentials \
--from-literal=GITHUB_USER=${GITHUB_USER} \
--from-literal=GITHUB_TOKEN=${GITHUB_TOKEN}

