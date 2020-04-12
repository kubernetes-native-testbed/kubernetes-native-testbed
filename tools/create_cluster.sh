#!/bin/bash
set -u

CURRENT_DIR=$(cd $(dirname $0); pwd)

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  gcloud beta container --project $GCP_PROJECT clusters create $CLUSTER_NAME --region $GCP_REGION --no-enable-basic-auth --release-channel "rapid" --machine-type "n1-standard-4" --image-type "UBUNTU" --disk-type "pd-ssd" --disk-size "50" --num-nodes "4" --enable-autoscaling --min-nodes "1" --max-nodes "15" --enable-autoupgrade --enable-autorepair --no-enable-stackdriver-kubernetes --addons HorizontalPodAutoscaling
  # --preemptible
  gcloud container clusters get-credentials $CLUSTER_NAME --region $GCP_REGION --project $GCP_PROJECT
fi
