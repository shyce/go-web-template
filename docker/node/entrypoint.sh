#!/bin/sh
if [ ! -f "package.json" ]; then
    if [ "$TYPESCRIPT" = true ] ; then
        npx create-react-app . --template typescript
    else
        npx create-react-app .
    fi
    jq '.proxy="http://golang:3000"' package.json | sponge package.json
else
    if [ ! -d "node_modules" ]; then
        yarn
    fi
fi

exec "$@"