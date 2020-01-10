#!/bin/bash
for SVC in $(ls microservices/); do
  kubectl --namespace ${SVC} apply -f manifests/${SVC}/ 
done

