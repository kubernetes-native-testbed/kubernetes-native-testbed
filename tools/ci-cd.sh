#!/bin/bash
sh docker-build.sh
sh manifests-check.sh
sh apply-manifests.sh

