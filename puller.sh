#!/bin/bash
pwd
cd $GOPATH/src/github.com/daflad/scrabble.nehpets.co.uk
while true; do
  git pull -q
  sleep 10
done
