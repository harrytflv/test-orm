#!/bin/bash

if [ "$(uname)" == "Darwin" ]; then
    PATH="$(brew --prefix coreutils)/libexec/gnubin:$PATH"
fi

GOMOD_ROOT=$(dirname $(dirname $(readlink -f $BASH_SOURCE)))
ROOT=$(dirname $(dirname $(readlink -f $BASH_SOURCE)))
COMPOSE_FILE=$ROOT/hack/docker-compose.yml

set -e

do_up() {
    docker-compose -f $COMPOSE_FILE -p foo up -d
}

do_down() {
    docker-compose -f $COMPOSE_FILE -p foo down
}

usage() {
    echo ""
    echo "Chong"
    echo ""
    echo "Usage: stack.sh [action]"
    echo ""
    echo "Actions:"
    echo "    up:        "
    echo "    down:      "
    echo ""
    exit 1
}

while true; do
    case $1 in
        up|down)
            COMMAND=$1
            shift
            set -x
            do_$COMMAND $@
            set +x
            break
            ;;
        ""|--help)
            usage ;;
        *)  
            echo "Bad command/argument: '$1'"
            exit 1
            ;;
    esac
done
