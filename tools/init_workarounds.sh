#!/bin/bash

# https://github.com/kubernetes-native-testbed/kubernetes-native-testbed/issues/65
until kubectl -n tekton-pipelines get deploy -l app.kubernetes.io/instance=ci-manifests-cd; do
  kubectl -n tekton-pipelines delete deployment el-github-listener
done


