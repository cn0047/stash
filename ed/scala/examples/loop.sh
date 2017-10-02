#!/bin/sh
exec scala "$0" "$@"
!#

val list = List(1,2,3,4,5);
for (l <- list) {
    println(l);
}
