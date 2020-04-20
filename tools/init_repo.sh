#!/bin/bash
set -eu

export LOADBALANCER_IP_ADDRESS=$(gcloud compute addresses list --filter="name=${LOADBALANCER_IP_NAME}" --format="value(address)")

CURRENT_DIR=$(cd $(dirname $0); pwd)

git checkout develop

for DIR in manifests microservices; do
  for FILE in `find ${CURRENT_DIR}/../${DIR}/ -type f`; do
    perl -pi -e "s|__LOADBALANCER_IP_ADDRESS__|${LOADBALANCER_IP_ADDRESS}|g" $FILE;
    perl -pi -e "s|__TB_GITHUB_ORG_NAME__|${TB_GITHUB_USER}|g" $FILE;
  done
done
perl -pi -e "s|__LOADBALANCER_IP_ADDRESS__|${LOADBALANCER_IP_ADDRESS}|g" ${CURRENT_DIR}/../README.md
perl -pi -e "s|__TB_GITHUB_ORG_NAME__|${TB_GITHUB_USER}|g" ${CURRENT_DIR}/../README.md

git add --all
git commit -m "initialized repo"
git push origin develop

