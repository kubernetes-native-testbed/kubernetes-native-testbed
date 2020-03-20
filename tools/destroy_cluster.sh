#!/bin/bash
if [ -z "$KUBERNETES_PLATFORM" ]; then
  echo "please set KUBERNETES_PLATFORM env var";
  exit 1;
fi

CURRENT_DIR=$(cd $(dirname $0); pwd)

kubectl -n argocd delete applications --all
kubectl -n projectcontour delete svc envoy

if [ $KUBERNETES_PLATFORM = "gke" ]; then
  sh ${CURRENT_DIR}/destroy_cluster_gke.sh
fi
