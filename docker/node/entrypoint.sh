#!/bin/sh
set -ex
if [ ! -f "package.json" ]; then
    if [ "$TYPESCRIPT" = true ] ; then
        npx create-react-app . --template typescript --scripts-version ${SCRIPTS_VERSION}
    else
        npx create-react-app . --scripts-version ${SCRIPTS_VERSION}
    fi
    jq ".proxy=\"http://golang:3000\"" package.json | sponge package.json
else
    if [ ! -d "node_modules" ]; then
        npm i
    fi
fi

exec "$@"