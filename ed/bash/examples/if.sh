#!/bin/bash

if [ $1 = '1' ]; then
    echo 'yes 1'
fi

if [ $1 = '2' ]; then
    echo 'yes 2'
else
    echo 'no else'
fi

if [ $1 = '3' ]; then
    echo 'yes 3'
elif [ $1 = '4' ]; then
    echo 'yes 4'
else
    echo 'no else else'
fi

# and
if [ $1 = 'a' -a $2 = 'b' ]; then
    echo 'ab = ✅'
else
    echo 'ab = ⛔️'
fi
