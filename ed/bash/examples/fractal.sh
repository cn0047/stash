#!/bin/bash

# Simple fractal implementation on bash.
#
# This script was written to solve task on hackerrank
# but maybe you'll find it interesting.
# @see https://monosnap.com/file/fDI4Ei0B5wN4YeFQhHfOMZh0Iy9Ext
# @see https://www.hackerrank.com/challenges/fractal-trees-all/problem

# Configurations:
rows=61             # Fractal image rows count.
columns=100         # Fractal image columns count.
patterheight=17     # Height of fractal patter (small part of whole fractal image).
c1='○'              # Custom character-1 (with purpose to configure it later).
c2='●'              # Custom character-2.
globalloops=15      # `read globalloops` # Global fractal patterns count (for hackerrank).
root=$((columns/2)) # Calculate column number for root node.

# Prints array $1 (which represents 1 line of fractal image) to stdout.
#
# @param array $1 Array to be printed.
function prn() {
  local -n arr=$1
  local len=${#arr[@]}
  local i=0
  for ((i=0; i<len; i++)); do
    echo -n ${arr[i]}
  done
  printf "\n"
}

# Generates blank fractal line (line which contains only character-1)
# with length equal to previous line length.
# This function intentionally receives array (not array length) as param
# because this approach looks more consistent with general code design.
#
# @param array $1 Array which represents previous fractal line.
# @category Pure function.
GetBlankLine() {
  local -n arr=$1
  local len=${#arr[@]}
  local i=0
  for ((i=0; i<len; i++)); do
    line[i]=${c1}
  done
  echo ${arr[@]}
}

# Generates fractal line which looks the same like provided into function one.
# This function created just for consistency of code design.
#
# @param array $1 Array which represents previous fractal line.
# @category Pure function.
GetRepeatedLine() {
  local -n arr=$1
  echo ${arr[@]}
}

# Generates branching (2 points instead 1) fractal line
# out of previous line (which is provided as param $1 into function).
# Draws 2 new points from both sides of previous point position.
#
# @param array $1 Array which represents previous fractal line.
# @category Pure function.
GetBranchingLine() {
  local -n arr=$1
  local len=${#arr[@]}
  local i=0
  for ((i=0; i<len; i++)); do
    if [ ${arr[i]} = $c2 ]; then
      arr[$((i-1))]=${c2}
      arr[$i]=${c1}
      arr[$((i+1))]=${c2}
      i=$((i+2))
    fi
  done
  echo ${arr[@]}
}

# Generates diverge fractal line out of previous line.
# This function depends on direction (which is continuously interchanging)
# and draws new point aside of previous point position.
#
# @param array $1 Array which represents previous fractal line.
# @category Pure function.
GetDivergeLine() {
  local -n arr=$1
  local len=${#arr[@]}
  local i=0
  local direction=1
  for ((i=0; i<len; i++)); do
    if [ ${arr[i]} = $c2 ]; then
      arr[$i]=${c1}
      if [[ $direction -eq 1 ]]; then
        direction=2
        arr[$((i-1))]=${c2}
      elif [[ $direction -eq 2 ]]; then
        direction=1
        arr[$((i+1))]=${c2}
        i=$((i+1))
      fi
    fi
  done
  echo ${arr[@]}
}

# Main function to run application and draw fractal image.
main() {
  # Initialize first fractal line (root).
  local line=()
  for ((i=0; i<columns; i++)); do
    line[i]=${c1}
    if [[ $i -eq $root ]]; then
      line[i]=${c2}
    fi
  done
  prn line

  local curheight=0 # Current fractal pattern height (just piece of whole fractal image).
  local stage=1     # Represents strategy to generate next fractal line.
  local gloop=1     # Global loop iteration id for global fractal patterns count.

  # Main loop.
  for ((i=1; i<rows; i++)) do
    curheight=$((curheight+1))

    # Determines stage depending on:
    # whether fractal pattern height is reached
    # or reached global fractal patterns count.
    if [[ $gloop -gt globalloops ]]; then
      stage=0
    elif [[ $curheight -eq $patterheight ]]; then
      curheight=0
      if [[ $stage -eq 1 ]]; then
        stage=2
      elif [[ $stage -eq 2 ]]; then
        stage=1
        patterheight=$((patterheight/2))
        gloop=$((gloop+1))
      fi
    fi

    # Call appropriate function to generate next fractal line
    # accordingly to stage (strategy).
    if [[ $stage -eq 0 ]]; then
      line=($(GetBlankLine line))
    elif [[ $stage -eq 1 ]]; then
      line=($(GetRepeatedLine line))
    elif [[ $stage -eq 2 && $curheight -eq 0 ]]; then
      line=($(GetBranchingLine line))
    elif [[ $stage -eq 2 ]]; then
      line=($(GetDivergeLine line))
    fi

    prn line
  done
}

# Run application:
main
