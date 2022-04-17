#!/usr/bin/env bash

. /Users/skeptycal/.dotfiles/zshrc_inc/functions/func_sys.zsh

set -e
echo "" >| coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go build -race -v .
    check_last "failed to compile"
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
