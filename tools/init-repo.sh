#!/bin/bash

LOADBALANCER_IP_ADDRESS=34.84.105.184

for FILE in `find . -type f`; do
  perl -pi -e "s|__LOADBALANCER_IP_ADDRESS__|${LOADBALANCER_IP_ADDRESS}|g" $FILE;
done



