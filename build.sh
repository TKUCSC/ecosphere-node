#!/bin/bash

if [[ $1 = "run" ]]
then
	echo "running..."

	./dist/node

	echo "done..."
elif [[ $1 = "build" ]]
then
	echo "building..."

	[ -e "./dist" ] && rm -rf ./dist
	cd src
	go build -o ../dist/node

	echo "done..."
else
	./build.sh build
	./build.sh run
fi
