#!/bin/bash

RUN_NAME="fabricServer"

mkdir -p output/bin
go build -o output/bin/${RUN_NAME}
