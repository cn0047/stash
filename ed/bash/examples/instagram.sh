<?php

# export csrf_token=$(curl https://www.instagram.com/ | grep -Po '(?<="csrf_token": ")[^"]*')

$cmd = <<<'NOWDOC'
export name="fcvk"`date +%s`
curl 'https://www.instagram.com/accounts/web_create_ajax/' \
-H 'Host: www.instagram.com' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:52.0) Gecko/20100101 Firefox/52.0' \
-H 'Accept: */*' \
-H 'Accept-Language: en-US,en;q=0.5' --compressed \
-H 'Referer: https://www.instagram.com/' \
-H 'X-CSRFToken: OYeX3gLVcALT73hmymeE9jT2yoxqs3g5' \
-H 'X-Instagram-AJAX: 1' \
-H 'Content-Type: application/x-www-form-urlencoded' \
-H 'X-Requested-With: XMLHttpRequest' \
-H 'Cookie: mid=WM0WdQAEAAGdcRKb3jAnjXYMHcI6; ig_vw=1359; ig_pr=1.8181818723678589; s_network=""; ig_dau_dismiss=1489924290094; csrftoken=OYeX3gLVcALT73hmymeE9jT2yoxqs3g5' \
-H 'Connection: keep-alive' --data 'email=codenamek2010%2B'$name'%40gmail.com&password=assdfdffghfy&username='$name'&first_name=full+'$name
echo $name
NOWDOC;
var_export(`$cmd`);
