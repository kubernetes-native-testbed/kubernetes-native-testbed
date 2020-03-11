#!/bin/bash
kubectl -n argocd delete applications --all
kubectl -n projectcontour delete svc envoy

