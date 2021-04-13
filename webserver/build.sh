#!/bin/bash

RUN_NAME="webServer"

mkdir -p output/bin
go build -o output/bin/${RUN_NAME}
