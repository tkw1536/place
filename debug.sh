#!/bin/bash

make
bin/place-server -debug=true -bind localhost:8080 -webhook /webhook/ -static $1 -script "$(pwd)/bin/place-git-update -from $2 -to $1 -ref refs/heads/master"