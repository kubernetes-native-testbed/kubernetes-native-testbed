#!/bin/bash
set -xeu

CURRENT_DIR=$(cd $(dirname $0); pwd)

source ${CURRENT_DIR}/env

sh ${CURRENT_DIR}/deallocate_staticip.sh

