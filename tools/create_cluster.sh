#!/bin/bash
set -u

CURRENT_DIR=$(cd $(dirname $0); pwd)

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  export GCP_SA_NAME=testbed-gcp-sa
  gcloud iam service-accounts create $GCP_SA_NAME

  export GCP_SA_EMAIL=$(gcloud iam service-accounts list --filter="Name:${GCP_SA_NAME}" --format='value(email)')

  gcloud beta container --project $GCP_PROJECT clusters create $CLUSTER_NAME --region $GCP_REGION --no-enable-basic-auth --release-channel "rapid" --machine-type "n1-standard-4" --image-type "UBUNTU" --disk-type "pd-ssd" --disk-size "50" --num-nodes "4" --enable-autoscaling --min-nodes "1" --max-nodes "15" --enable-autoupgrade --enable-autorepair --no-enable-stackdriver-kubernetes --addons HorizontalPodAutoscaling --scopes="" --service-account=$GCP_SA_EMAIL
  # --preemptible
  gcloud container clusters get-credentials $CLUSTER_NAME --region $GCP_REGION --project $GCP_PROJECT
fi
