#!/bin/bash

# echo 'https://mediaservice.audi.com/media/live/50900/fly1400x601n8/8w2/2018.png?wid=291' \
#   | ed/bash/examples/pipe.ImagesHtml.sh > /tmp/x.html && open /tmp/x.html

cat << HTML
<html>
<head>
  <style>
    #root {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
    }
    img {
      height: 200px;
      margin: 1px;
      border: 1px solid gainsboro;
    }
  </style>
</head>
<body><div id="root">
HTML

while IFS= read -r imgSrc; do
  if [[ $imgSrc ]]; then
    echo  '<img src="'${imgSrc}'" >'
  fi
done

echo '</div></body></html>'
