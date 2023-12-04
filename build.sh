#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o promonitor
chmod +x promonitor
mv -f promonitor ./scripts/