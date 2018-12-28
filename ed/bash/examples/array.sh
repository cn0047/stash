#!/bin/bash

declare -a arr1=("element1" "element2" "element3")

declare -A arr2
arr2['first']='First element'
arr2['second']='Second element'

declare -A arr3=(["a"]="x" ["b"]="q" )

echo "${arr1[1]}"
printf "\n"

for i in "${arr1[@]}"; do
   echo "$i"
done

printf "\n"
echo "${arr2[@]}"

files=( * )
printf "\nFiles:\n"
for file in "${files[@]}"; do
  echo "$file"
done
