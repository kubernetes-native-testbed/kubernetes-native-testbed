#!/bin/bash

KREW_PLUGINS=(sort-manifests)

for KREW_PLUGIN in $KREW_PLUGINS; do
  kubectl krew install $KREW_PLUGIN
done

