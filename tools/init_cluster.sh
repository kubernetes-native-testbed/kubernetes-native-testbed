#!/bin/bash
CURRENT_DIR=$(cd $(dirname $0); pwd)
kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/ | kubectl apply -f -

kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/ | kubectl apply -f -

kubectl delete -f ${CURRENT_DIR}/../manifests/infra/tmp -R

