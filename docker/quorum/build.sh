#!/bin/bash

VERSION="1.0"
docker build -t test-node:${VERSION} -t test-node:latest .
