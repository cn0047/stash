#!/bin/bash

v=`echo -n "my msg" | base64`
m=`echo -n $v | base64 -D`
printf "in   base64: $v \n"
printf "from base64: $m \n"
