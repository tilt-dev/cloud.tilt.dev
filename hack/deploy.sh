#!/bin/bash
# Deploys cloud.tilt.dev

cd $(dirname $(dirname $0))

set -ex

./hack/build.sh

TODAY=$(date +"%Y-%m-%d")
SECONDS=$(date +"%s")
TAG="$TODAY-$SECONDS"
docker build -t "gcr.io/windmill-prod/cloud-tilt-dev:$TAG" .
docker push "gcr.io/windmill-prod/cloud-tilt-dev:$TAG"

helm upgrade --install cloud-tilt-dev ./chart \
     --set "imageName=gcr.io/windmill-prod/cloud-tilt-dev:$TAG"
