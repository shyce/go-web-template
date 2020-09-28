#!/bin/sh
if [ ! -f "package.json" ]; then
    if [ "$TYPESCRIPT" = true ] ; then
        npx create-react-app . --template typescript
    else
        npx create-react-app .
    fi
else
    npm i
fi

exec "$@"