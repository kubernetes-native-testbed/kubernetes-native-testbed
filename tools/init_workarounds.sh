#!/bin/bash

# https://github.com/kubernetes-native-testbed/kubernetes-native-testbed/issues/65
until kubectl -n tekton-pipelines get deploy el-github-listener -o jsonpath="{.metadata.labels.app\.kubernetes\.io/instance}" --allow-missing-template-keys=false; do
  kubectl -n tekton-pipelines delete deployment el-github-listener
done


