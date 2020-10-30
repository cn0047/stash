#!/bin/bash

# @example:
# grabDockerHubTags 'tensorflow/tensorflow'
# grabDockerHubTags 'cn007b/ubuntu'
grabDockerHubTags() {
  repo=$1
  rm -rf ./dockerHubTags && mkdir ./dockerHubTags && cd ./dockerHubTags
  ua='User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36'
  i=0
  while true; do
    i=$(($i+1))
    echo -n "\rProcessing page: $i"
    file="?page_size=100&page=$i"
    curl -s -O -H $ua "https://hub.docker.com/v2/repositories/$repo/tags/$file"
    cat $file | jq -r '.results[].name' >> tags
    next=`cat $file | jq .next`
    rm -rf $file
    if [[ $next = 'null' ]]; then break; fi
  done
  echo "\nFound tags:"
  cat tags | sort
  cd ..
}
