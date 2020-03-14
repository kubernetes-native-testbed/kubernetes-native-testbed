#!/bin/bash

# install skaffold
SKAFFOLD_VERSION=1.5.0
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v${SKAFFOLD_VERSION}/skaffold-linux-amd64
chmod +x skaffold
sudo mv skaffold /usr/local/bin

# install telepresence
brew cask install osxfuse
brew install datawire/blackbird/telepresence

