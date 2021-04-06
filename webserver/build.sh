#!/bin/bash

RUN_NAME="cloudStore"

mkdir -p output/bin
go build -o output/bin/${RUN_NAME}
