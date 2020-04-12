#!/bin/bash
set -u
CURRENT_DIR=$(cd $(dirname $0); pwd)

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  gcloud container clusters delete ${CLUSTER_NAME} --region ${GCP_REGION} --quiet

  export GCP_SA_NAME=testbed-gcp-sa
  export GCP_SA_EMAIL=$(gcloud iam service-accounts list --filter="Name:${GCP_SA_NAME}" --format='value(email)')

  gcloud iam service-accounts delete $GCP_SA_EMAIL --quiet
fi
