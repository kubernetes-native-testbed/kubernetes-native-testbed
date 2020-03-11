#!/bin/bash

if [ -z "$PROJECT" ]; then
  echo "please set PROJECT env var";
  exit 1;
fi
export CLUSTER=tb-$RANDOM
export REGION=asia-northeast1

gcloud beta container --project $PROJECT clusters create $CLUSTER --region $REGION --no-enable-basic-auth --release-channel "stable" --machine-type "n1-standard-4" --image-type "UBUNTU" --disk-type "pd-ssd" --disk-size "150" --num-nodes "2" --enable-autoscaling --min-nodes "1" --max-nodes "10" --enable-autoupgrade --enable-autorepair

gcloud container clusters get-credentials $CLUSTER --region $REGION --project $PROJECT


