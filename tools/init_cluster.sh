#!/bin/bash
set +e

CURRENT_DIR=$(cd $(dirname $0); pwd)
until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/ | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

cat ${CURRENT_DIR}/../manifests/infra/tmp/* | kubectl delete -f -
set -e

