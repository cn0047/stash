#!/bin/bash
deploy=false
uglify=false
while (( $# > 1 )); do case $1 in
    --deploy) deploy="$2";;
    --uglify) uglify="$2";;
    *) break;
  esac; shift 2
done
$deploy && echo "will deploy... deploy = $deploy"
$uglify && echo "will uglify... uglify = $uglify"

foo() {
  while [[ "$#" -gt 0 ]]
  do
    case $1 in
      -f|--follow)
        local FOLLOW="following"
        ;;
      -t|--tail)
        local TAIL="tail=$2"
        ;;
    esac
    shift
  done
  echo "FOLLOW: $FOLLOW"
  echo "TAIL: $TAIL"
}
# foo -f
# foo -t 10
# foo -f --tail 10
# foo --follow --tail 10
