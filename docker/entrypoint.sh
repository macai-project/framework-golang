#!/bin/bash
set -e

echo -n "Copying personal certificates into root folder... "
cp /.ssh /root/.ssh -R
echo -n "fixing certificates permission... "
chmod 700 /root/.ssh
chmod 400 /root/.ssh/*
echo "done!"

echo "Installing gotest and wait-for and Delve"
go install github.com/rakyll/gotest@latest
go install github.com/alioygur/wait-for@latest
go install github.com/radovskyb/watcher/...@latest

if [ -n "$1" ]; then
  echo "Launching $1..."
  exec "$@"
else
  echo "Launching Tests..."
  # Watch src directory and re-build app
  watcher -cmd="dlv debug --build-flags='github.com/macai-project/framework-golang/pkg/framework' --headless --listen=:2345 --api-version=2 --accept-multiclient" -dotfiles=false --keepalive -startcmd=true
fi