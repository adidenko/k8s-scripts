#!/bin/bash

for i in {0..9} ; do
  size=$((200 + $RANDOM % 200))
  files=$((800 + $RANDOM % 400))
  echo
  echo
  echo "Generating 500 directories, $files files in each, ${size}K each file"
  ./mdbench.py -n -r -d 500 -f $files -s ${size}K --chunk 100K --direct /drives/data61/test_gen_payload_01
done
