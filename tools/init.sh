#!/bin/bash
set -xeu

CURRENT_DIR=$(cd $(dirname $0); pwd)

source ${CURRENT_DIR}/env

sh ${CURRENT_DIR}/init_cli.sh
sh ${CURRENT_DIR}/allocate_staticip.sh
sh ${CURRENT_DIR}/init_repo.sh

