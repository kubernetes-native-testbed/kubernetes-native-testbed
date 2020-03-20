#!/bin/bash
set +e

CURRENT_DIR=$(cd $(dirname $0); pwd)
kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/ | kubectl apply -f -

kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/ | kubectl apply -f -

cat ${CURRENT_DIR}/../manifests/infra/tmp/* | kubectl delete -f -
set -e

