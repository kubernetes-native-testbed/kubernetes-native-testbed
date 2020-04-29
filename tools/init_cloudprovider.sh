#!/bin/bash
set -u

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  gcloud services enable container.googleapis.com --project ${GCP_PROJECT}
fi
