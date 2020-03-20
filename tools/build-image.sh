#!/bin/bash
CURRENT_DIR=$(cd $(dirname $0); pwd)
cd $CURRENT_DIR

for DIR in `find images/* -type f -name Dockerfile | sed 's|Dockerfile||'`; do
  IMAGENAME=k8stestbed/`echo $DIR | awk -F '/' '{$1="";print $0}' | sed -e 's| ||g'`;
  docker build $DIR -t $IMAGENAME;
  docker push $IMAGENAME;
done

cd -
