#!/usr/bin/env bash

go install github.com/radovskyb/watcher/...@latest
go install github.com/rakyll/gotest@latest
watcher -cmd="gotest -race -count=1 -v -coverpkg=./... -coverprofile=coverage.out ./..." -dotfiles=false --keepalive -startcmd=true -ignore=coverage.out