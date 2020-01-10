#!/bin/bash
for SVC in $(ls microservices/); do
  conftest test --policy tools/policy ./manifests/${SVC}/*.yaml
done

