#!/bin/bash

VERSION="1.0"
docker build -t test-base:${VERSION} -t test-base:latest .
