FROM --platform=linux/amd64 ubuntu:20.04

##### basic install #####
RUN apt-get update && apt-get install -y --no-install-recommends apt-utils
RUN apt-get install -y \
    curl \
    git \
    wget \
    apt-transport-https \
    vim \
    net-tools \
    build-essential \
    jq

CMD ["tail", "-f", "/dev/null"]
