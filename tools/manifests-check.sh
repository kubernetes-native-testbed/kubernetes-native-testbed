#!/bin/bash

CURRENT_DIR=$(cd $(dirname $0); pwd)

for MICROSERVICE in `find ${CURRENT_DIR}/../microservices -name Dockerfile | awk -F "/" '{print $(NF-1)}'`; do
  conftest test --policy ${CURRENT_DIR}/policy ${CURRENT_DIR}/../manifests/${MICROSERVICE}/*.yaml
done

