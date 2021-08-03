#!/usr/bin/env bash

go get -u github.com/radovskyb/watcher/...
go get -u github.com/rakyll/gotest
watcher -cmd="gotest -race -count=1 -v -coverpkg=./... -coverprofile=coverage.out ./..." -dotfiles=false --keepalive -startcmd=true -ignore=coverage.out