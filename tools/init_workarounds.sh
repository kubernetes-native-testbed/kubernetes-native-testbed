#!/bin/bash

# https://github.com/kubernetes-native-testbed/kubernetes-native-testbed/issues/65
kubectl -n tekton-pipelines delete deployment el-github-listener


