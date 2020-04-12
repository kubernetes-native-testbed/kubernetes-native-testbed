#!/bin/bash
set -u

CURRENT_DIR=$(cd $(dirname $0); pwd)

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  gcloud compute addresses delete ${LOADBALANCER_IP_NAME} --project=${GCP_PROJECT} --region=${GCP_REGION} --quiet
fi
