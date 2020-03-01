#!/bin/bash

export CLUSTER=tb1
export REGION=asia-northeast1

gcloud beta container --project $PROJECT clusters create $CLUSTER --region $REGION --no-enable-basic-auth --release-channel "rapid" --machine-type "n1-standard-4" --image-type "UBUNTU" --disk-type "pd-ssd" --disk-size "150" --num-nodes "2" --enable-autoscaling --min-nodes "1" --max-nodes "10" --enable-autoupgrade --enable-autorepair

gcloud container clusters get-credentials $CLUSTER --region $REGION --project $PROJECT


