#!/bin/bash
CURRENT_DIR=$(cd $(dirname $0); pwd)

sh ${CURRENT_DIR}/init_cli.sh
sh ${CURRENT_DIR}/create_cluster.sh
sh ${CURRENT_DIR}/init_cluster.sh
