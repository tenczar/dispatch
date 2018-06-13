#!/bin/bash

### COMMON VARIABLES ###
: ${DOCKER_REGISTRY:="dispatchframework"}
: ${BASE_IMAGE_PYTHON3:="python3-base:0.0.8"}
: ${BASE_IMAGE_NODEJS6:="nodejs-base:0.0.8"}
: ${BASE_IMAGE_POWERSHELL:="powershell-base:0.0.9"}
: ${BASE_IMAGE_JAVA:="java-base:0.0.9"}
: ${IMAGES_YAML:="images.yaml"}
