#!/bin/bash

f() {
    return $(( $1 + $1))
}

f2() {
    echo $(( $1 + $1))
}

startsWithA() {
    [[ $1 == [Aa]* ]];
}

f 2
echo $? # result: 4
echo $(f2 3) # result: 6
if startsWithA "apple"; then echo "startsWithA - yes"; else echo "startsWithA - no"; fi # yes
