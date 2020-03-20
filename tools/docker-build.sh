#!/bin/bash
PUSH_ENABLED=${PUSH_ENABLED:-true}
TAG=${TAG:-latest}

for SVC in $(ls microservices/); do
  docker build ./microservices/${SVC} -t k8stestbed/${SVC}:${TAG}; 
  if "${PUSH_ENABLED}"; then
    docker push k8stestbed/${SVC}:${TAG};
  fi
done

