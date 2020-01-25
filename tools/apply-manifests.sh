#!/bin/bash
for SVC in $(ls manifests/); do
  kubectl  --validate=false --namespace ${SVC} apply -f manifests/${SVC}/ -R 
done

