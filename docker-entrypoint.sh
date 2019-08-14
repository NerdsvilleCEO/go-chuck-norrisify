#!/bin/bash

if [ ! -z "$ENVIRONMENT" -a "$ENVIRONMENT" == 'development' ]; then
    go get github.com/codegangsta/gin
    go-wrapper download
    go-wrapper install
else
    echo "Nothing to do"
fi

exec "$@"
