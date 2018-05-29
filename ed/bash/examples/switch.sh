#!/bin/bash

# RUN:
# . ed/bash/examples/switch.sh one
# . ed/bash/examples/switch.sh 3

case $1 in
one)
  echo "run command 'one'"
  ;;
two)
  echo "run command 'two'"
  ;;
*)
  echo "run command 'default'"
  ;;
esac
