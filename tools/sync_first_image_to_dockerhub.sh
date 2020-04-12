#!/bin/bash

FROM_IMAGE_REPO=${FROM_IMAGE_REPO:-harbor.xxx.xxx.xxx.xxx.nip.io}
FROM_IMAGE_TAG=${FROM_IMAGE_TAG:-c57ea1d59ad621a4cd3bdae8d334dd83952cfbd5}

CURRENT_DIR=$(cd $(dirname $0); pwd)

for MICROSERVICE in `find ${CURRENT_DIR}/../microservices -name Dockerfile | awk -F "/" '{print $(NF-1)}'`; do
  echo "===== syncing $MICROSERVICE image ====================================================="
  docker pull ${FROM_IMAGE_REPO}/library/${MICROSERVICE}:${FROM_IMAGE_TAG}
  docker tag ${FROM_IMAGE_REPO}/library/${MICROSERVICE}:${FROM_IMAGE_TAG} k8stestbed/$MICROSERVICE:latest-cache 
  docker push k8stestbed/$MICROSERVICE:latest-cache
done

