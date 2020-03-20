#!/bin/bash

CURRENT_DIR=$(cd $(dirname $0); pwd)

if [ -z "$KUBERNETES_PLATFORM" ]; then
  echo "please set KUBERNETES_PLATFORM env var";
  exit 1;
fi

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  sh ${CURRENT_DIR}/create_cluster_gke.sh
fi
