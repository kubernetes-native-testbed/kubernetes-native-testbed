#!/bin/bash
set -xeu

CURRENT_DIR=$(cd $(dirname $0); pwd)

source ${CURRENT_DIR}/env

kubectl -n argocd delete applications --all
kubectl -n projectcontour delete svc envoy

${CURRENT_DIR}/destroy_cluster.sh

