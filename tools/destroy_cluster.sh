#!/bin/bash
set -u
CURRENT_DIR=$(cd $(dirname $0); pwd)

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  gcloud container clusters delete ${CLUSTER_NAME} --region ${GCP_REGION} --quiet
fi
