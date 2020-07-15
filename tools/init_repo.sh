#!/bin/bash
set -eu

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  export LOADBALANCER_IP_ADDRESS=$(gcloud compute addresses list --filter="name=${LOADBALANCER_IP_NAME}" --format="value(address)")
fi

echo ${TB_GITHUB_ORG_NAME:?Variable not set. Aborting...}
echo ${LOADBALANCER_IP_ADDRESS:?Variable not set. Aborting...}

echo "Repository will be initialized with the following information."
echo "TB_GITHUB_ORG_NAME=${TB_GITHUB_ORG_NAME}"
echo "KUBERNETES_PLATFORM=${KUBERNETES_PLATFORM}"
echo "LOADBALANCER_IP_ADDRESS=${LOADBALANCER_IP_ADDRESS}"
echo "Would you like to proceeed? (y/n)"

read param

case "$param" in
  [Yy])
    echo 'Initialization starting...'
    ;;
  [Nn])
    echo "Initialization cancelling..."
    exit 0
    ;;
  *)
    echo "Invalid parameter. Please type y or n!"
    exit 1
    ;;
esac

CURRENT_DIR=$(cd $(dirname $0); pwd)

git checkout develop

for DIR in manifests microservices; do
  for FILE in `find ${CURRENT_DIR}/../${DIR}/ -type f`; do
    perl -pi -e "s|__LOADBALANCER_IP_ADDRESS__|${LOADBALANCER_IP_ADDRESS}|g" $FILE;
    perl -pi -e "s|__TB_GITHUB_ORG_NAME__|${TB_GITHUB_ORG_NAME}|g" $FILE;
  done
done
perl -pi -e "s|__LOADBALANCER_IP_ADDRESS__|${LOADBALANCER_IP_ADDRESS}|g" ${CURRENT_DIR}/../README.md
perl -pi -e "s|__TB_GITHUB_ORG_NAME__|${TB_GITHUB_ORG_NAME}|g" ${CURRENT_DIR}/../README.md

git add --all
git commit -m "initialized repo"
git push origin develop

