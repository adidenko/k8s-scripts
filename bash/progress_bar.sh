#!/bin/bash

# Usage: Call the function with the current progress and total steps.
function progress_bar {
  local PROGRESS=$1
  local TOTAL=$2
  local BAR_LENGTH=50
  local FILLED_LENGTH=$((BAR_LENGTH * PROGRESS / TOTAL))
  local EMPTY_LENGTH=$((BAR_LENGTH - FILLED_LENGTH - 1))

  printf "\r["
  printf "%0.s#" $(seq 0 $FILLED_LENGTH)
  if [ $PROGRESS -lt $TOTAL ]; then
    printf "%0.s-" $(seq 0 $EMPTY_LENGTH)
  fi
  printf "] %d%%" $((PROGRESS * 100 / TOTAL))

  if [ $PROGRESS -ge $TOTAL ]; then
    echo -e "\nDone!"
  fi
}

# Simulate a process with random sleep times
echo "Test progress bar"
MAX=30
for i in $(seq 0 $MAX); do
  sleep $(awk -v min=0.1 -v max=0.5 'BEGIN{srand(); print min+rand()*(max-min)}')
  progress_bar $i $MAX
done
