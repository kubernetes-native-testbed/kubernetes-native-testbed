#!/bin/bash

OS_TYPE=$(uname | tr '[A-Z]' '[a-z]')

# install skaffold
if ! which skaffold; then
  SKAFFOLD_VERSION=1.5.0
  curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v${SKAFFOLD_VERSION}/skaffold-${OS_TYPE}-amd64
  chmod +x skaffold
  sudo mv skaffold /usr/local/bin
fi

# install telepresence
if ! which telepresence; then
  brew cask install osxfuse
  brew install datawire/blackbird/telepresence
fi

if [ -z "${REMOTE_CONTEXT}" ]; then
  echo "Please input REMOTE_CONTEXT env var"
  read REMOTE_CONTEXT
  export REMOTE_CONTEXT
fi
if [ -z "${LOCAL_CONTEXT}" ]; then
  echo "Please input LOCAL_CONTEXT env var"
  read LOCAL_CONTEXT
  export LOCAL_CONTEXT
fi

