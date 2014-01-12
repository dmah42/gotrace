#!/bin/bash

export GOPATH=`pwd`

go test trace && \
  go install trace && \
  go build testbed
