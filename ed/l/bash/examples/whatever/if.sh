#!/bin/bash

s='1'
if [ $s = '1' ]; then
    echo 'yes 1'
fi

if [ $s = '2' ]; then
    echo 'yes 2'
else
    echo 'no else'
fi

s='4'
if [ $s = '3' ]; then
    echo 'yes 3'
elif [ $s = '4' ]; then
    echo 'yes 4'
else
    echo 'no else'
fi

# and
a='a'
b='b'
if [ $a = 'a' -a $b = 'b' ]; then
    echo 'ab = ✅'
else
    echo 'ab = ⛔️'
fi

# or
a='0'
b='1'
if [ $a = '1' ] || [ $b = '1' ]; then
    echo 'yes'
else
    echo 'no'
fi
