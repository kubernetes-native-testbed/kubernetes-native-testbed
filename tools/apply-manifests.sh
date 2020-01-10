#!/bin/bash
for SVC in $(ls manifests/); do
  kubectl --namespace ${SVC} apply -f manifests/${SVC}/ -R 
done

