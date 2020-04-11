#!/bin/bash
export LOADBALANCER_IP_ADDRESS=$(gcloud compute addresses list --filter="name=testbed-ip" --format="value(address)")
export GITHUB_ORG_NAME=$(git config --get remote.origin.url | cut -d ":" -f 2 | cut -d "/" -f 1)

CURRENT_DIR=$(cd $(dirname $0); pwd)

git checkout develop

for FILE in `find . -type f`; do
  perl -pi -e "s|__LOADBALANCER_IP_ADDRESS__|${LOADBALANCER_IP_ADDRESS}|g" $FILE;
  perl -pi -e "s|__GITHUB_ORG_NAME__|${GITHUB_ORG_NAME}|g" $FILE;
done

git add --all
git commit -m "initialized repo"
git push origin develop

