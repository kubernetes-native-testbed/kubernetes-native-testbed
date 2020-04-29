#!/bin/bash
set -xeu

CURRENT_DIR=$(cd $(dirname $0); pwd)

source ${CURRENT_DIR}/env

${CURRENT_DIR}/init_cli.sh
${CURRENT_DIR}/enable_service.sh
${CURRENT_DIR}/allocate_staticip.sh
${CURRENT_DIR}/init_repo.sh
