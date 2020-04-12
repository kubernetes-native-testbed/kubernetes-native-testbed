#!/bin/bash
set +e

CURRENT_DIR=$(cd $(dirname $0); pwd)

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/ns | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/system | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/infra/contour | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/infra/rook | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/infra | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/cicd/ci-manifests | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/cicd/cd-manifests/microservices | kubectl apply -f -; do
  echo retrying apply manifests at first time;
  sleep 1;
done

for MICROSERVICE in `find ${CURRENT_DIR}/../microservices -name Dockerfile | awk -F "/" '{print $(NF-1)}'`; do
  echo "applying ${MICROSERVICE} manifests"
  until kubectl sort-manifests -R -f ${CURRENT_DIR}/../manifests/${MICROSERVICE} | kubectl apply -f -; do
    echo retrying apply manifests at first time;
    sleep 1;
  done
done

cat ${CURRENT_DIR}/../manifests/infra/tmp/* | kubectl delete -f -
set -e

