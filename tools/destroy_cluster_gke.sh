#!/bin/bash
if [ -z "$GCP_PROJECT" ]; then
  echo "please set GCP_PROJECT env var";
  exit 1;
fi
export CLUSTER=${CLUSTER:-tb-$RANDOM}
export GCP_REGION=${GCP_REGION:-asia-northeast1}

# TODO
# glouc destroy cluster...

