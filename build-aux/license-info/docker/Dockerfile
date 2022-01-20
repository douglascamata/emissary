# syntax = docker/dockerfile:1.3-labs
### This file generates the list of dependencies used by an npm application

FROM node:10-alpine as dependencies

WORKDIR /app
ADD . .

RUN apk add --update --no-cache jq

RUN <<EOF sh -ex
    mkdir /dependencies
    npm install

    npm install -g license-checker@25.0.1
EOF

RUN <<EOF
    set -ex
    echo '{
            "name": "",
            "version": "",
            "description": "",
            "licenseFile": "",
            "licenseText": "",
            "licenseModified": ""
          }' > /app/customLicenseFormat.json

    export THIS_PACKAGE=$(cat package.json | jq -r '.name + "@" + .version')
    license-checker --excludePackages $THIS_PACKAGE --customPath /app/customLicenseFormat.json --json > /dependencies/dependencies.json
EOF

FROM scratch AS export-stage
COPY --from=dependencies /dependencies/dependencies.json .