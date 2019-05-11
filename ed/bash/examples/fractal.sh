#!/bin/bash

# Simple fractal implementation on bash.
#
# @see https://monosnap.com/file/kQ8EMJXEGsRxXIl8UfmsLznoaQAXmc

rows=63
columns=100
height=16
c1='○'
c2='●'

globalLoops=5 # read globalLoops # for hackerrank.
root=$((columns/2))

function prn() {
  local -n arr=$1
  local len=${#arr[@]}
  local i=0
  for ((i=0; i<len; i++)); do
      echo -n ${arr[i]}
  done
  printf "\n"
}

getRootLine() {
  local -n arr=$1
  local len=${#arr[@]}
  local i=0
  for ((i=0; i<len; i++)); do
    line[i]=${c1}
  done
  echo ${arr[@]}
}

getRepeatedLine() {
  local -n arr=$1
  echo ${arr[@]}
}

getBranchingLine() {
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

getDiagonalLine() {
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

# Initialize first (root) line.
line=()
for ((i=0; i<columns; i++)); do
  line[i]=${c1}
  if [[ $i -eq $root ]]; then
    line[i]=${c2}
  fi
done
prn line

curheight=0
stage=1
gloop=1
# Main loop.
for ((i=1; i<rows; i++)) do
  curheight=$((curheight+1))
  if [[ $curheight -eq $height ]]; then
    curheight=0
    if [[ $stage -eq 1 ]]; then
      stage=2
    elif [[ $stage -eq 2 ]]; then
      stage=1
      height=$((height/2))
      gloop=$((gloop+1))
    fi
  fi
  if [[ $gloop -gt globalLoops ]]; then
    stage=0
  fi

  if [[ $stage -eq 0 ]]; then
    line=($(getRootLine line))
  elif [[ $stage -eq 1 ]]; then
    line=($(getRepeatedLine line))
  elif [[ $stage -eq 2 && $curheight -eq 0 ]]; then
    line=($(getBranchingLine line))
  elif [[ $stage -eq 2 ]]; then
    line=($(getDiagonalLine line))
  fi

  prn line
done
