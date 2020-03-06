#!/bin/bash
kubectl sort-manifests -R -f ../manifests/ | kubectl apply -f -

