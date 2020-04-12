#!/bin/bash

kubectl -n tekton-pipelines create secret generic github-credentials \
--from-literal=TB_GITHUB_USER=${TB_GITHUB_USER} \
--from-literal=TB_GITHUB_TOKEN=${TB_GITHUB_TOKEN}

