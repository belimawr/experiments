#!/bin/bash
docker run --rm -v "$PWD":/usr/src/GoCD-Build -w /usr/src/GoCD-Build golang:1.7 make -i
