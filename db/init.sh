#!/bin/bash
#TODO Clean redis
for filename in data/*.json; do
    if [ -f "$filename" ]; then
        mongoimport --db "$1" --jsonArray --drop --file "$filename"
    fi
done
