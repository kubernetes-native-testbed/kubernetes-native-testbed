#!/bin/bash
SCRIPT_DIR=`dirname $0`
sh ${SCRIPT_DIR}/docker-build.sh
sh ${SCRIPT_DIR}/manifests-check.sh
sh ${SCRIPT_DIR}/apply-manifests.sh

