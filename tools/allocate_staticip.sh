#!/bin/bash

CURRENT_DIR=$(cd $(dirname $0); pwd)
source ${CURRENT_DIR}/env

if [ -z "$KUBERNETES_PLATFORM" ]; then
  echo "please set KUBERNETES_PLATFORM env var";
  exit 1;
fi

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  gcloud compute addresses create testbed-ip --project=${GCP_PROJECT} --region=${GCP_REGION}
  export LOADBALANCER_IP_ADDRESS=$(gcloud compute addresses list --filter="name=testbed-ip" --format="value(address)")
  echo "Assigned IP Address is ${LOADBALANCER_IP_ADDRESS}"
fi
