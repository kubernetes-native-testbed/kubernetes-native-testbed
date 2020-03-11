#!/bin/bash
for STORAGECLASS in `kubectl get storageclass -o jsonpath="{.items[*].metadata.name}"`; do
  kubectl patch storageclass $STORAGECLASS -p '{"metadata": {"annotations": {"storageclass.kubernetes.io/is-default-class": "false"}}}'
done

