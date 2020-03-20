#!/bin/bash
set -xeu

CURRENT_DIR=$(cd $(dirname $0); pwd)

source ${CURRENT_DIR}/env
sh ${CURRENT_DIR}/init_cli.sh
sh ${CURRENT_DIR}/create_cluster.sh
sh ${CURRENT_DIR}/init_clean_cloudprovider_config.sh
sh ${CURRENT_DIR}/init_cluster.sh
sh ${CURRENT_DIR}/init_credential.sh
sh ${CURRENT_DIR}/init_workarounds.sh

