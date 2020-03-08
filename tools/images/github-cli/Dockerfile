FROM ubuntu:18.04

RUN apt-get -y update && \
    apt-get -y install git sed curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN GITHUB_CLI_VERSION=0.6.1 && \
    curl -sL -o /root/gh.deb https://github.com/cli/cli/releases/download/v${GITHUB_CLI_VERSION}/gh_${GITHUB_CLI_VERSION}_linux_amd64.deb && \
    dpkg -i /root/gh.deb && \
    rm /root/gh.deb

