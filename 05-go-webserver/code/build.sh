#!/bin/bash

docker image build -t bogd/go-webserver .
docker container run bogd/go-webserver


