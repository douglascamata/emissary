#!/bin/env bash
set -ex

SCRIPT_FOLDER=$(pwd)
DOCKERFILE=$(pwd)/docker

pushd /home/andres/source/production/emissary/tools/src/js-mkopensource
go build -o "${SCRIPT_FOLDER}/js-mkopensource" .
popd

pushd /home/andres/source/production/emissary

find -name package.json -exec dirname {} \; | while IFS=$'\n' read packagedir; do
  pushd "${packagedir}"
  DOCKER_BUILDKIT=1 docker build -f "${DOCKERFILE}/Dockerfile" --output ${SCRIPT_FOLDER}/out/ .
  popd

  FILENAME="$(date +%s_%N).json"
  mv ${SCRIPT_FOLDER}/out/dependencies.json "${SCRIPT_FOLDER}/${FILENAME}"

  cat "${SCRIPT_FOLDER}/${FILENAME}" | "${SCRIPT_FOLDER}/js-mkopensource" | jq
done

popd