#!/bin/bash
set -eu

export LOADBALANCER_IP_ADDRESS=$(gcloud compute addresses list --filter="name=${LOADBALANCER_IP_NAME}" --format="value(address)")

CURRENT_DIR=$(cd $(dirname $0); pwd)

git checkout develop

for DIR in manifests microservices; do
  for FILE in `find ../${DIR} -type f`; do
    perl -pi -e "s|34.84.31.199|${LOADBALANCER_IP_ADDRESS}|g" $FILE;
    perl -pi -e "s|kubernetes-native-testbed-user|${TB_GITHUB_ORG_NAME}|g" $FILE;
  done
done

git add --all
git commit -m "initialized repo"
git push origin develop

