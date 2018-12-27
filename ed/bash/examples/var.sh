#!/usr/bin/env bash

x=204
echo x # Use of a variable like $X instead of ${X}
echo ${x}

echo "$0"
# echo "$(readlink -f "$0")"
