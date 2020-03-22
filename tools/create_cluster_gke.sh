#!/bin/bash
if [ -z "$GCP_PROJECT" ]; then
  echo "please set GCP_PROJECT env var";
  exit 1;
fi
export CLUSTER=${CLUSTER:-tb-$RANDOM}
export GCP_REGION=${GCP_REGION:-asia-northeast1}

gcloud beta container --project $GCP_PROJECT clusters create $CLUSTER --region $GCP_REGION --no-enable-basic-auth --release-channel "rapid" --machine-type "n1-standard-4" --image-type "UBUNTU" --disk-type "pd-ssd" --disk-size "50" --num-nodes "3" --enable-autoscaling --min-nodes "1" --max-nodes "15" --enable-autoupgrade --enable-autorepair --no-enable-stackdriver-kubernetes --addons HorizontalPodAutoscaling
# --preemptible

gcloud container clusters get-credentials $CLUSTER --region $GCP_REGION --project $GCP_PROJECT


