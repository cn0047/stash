#!/bin/bash


if [ -z "$1" ]; then
    echo 'Please specify action.'
    exit 1
fi
action=$1

create () {
    export name="fcvk"`date +%s`
    printf "\n C $name "
    curl -s 'https://www.instagram.com/accounts/web_create_ajax/' \
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
}

create2 () {

    curl 'https://www.instagram.com/accounts/web_create_ajax/attempt/' \
    -H 'Host: www.instagram.com' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:52.0) Gecko/20100101 Firefox/52.0' \
    -H 'Accept: */*' \
    -H 'Accept-Language: en-US,en;q=0.5' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'X-CSRFToken: DQds9PuX41a5BiuewGHDZeSAuzeIZFWu' \
    -H 'X-Instagram-AJAX: 1' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -H 'Referer: https://www.instagram.com/' \
    -H 'Cookie: mid=WOUeqAAEAAFrY7VM1GmjUKbEwxXL; csrftoken=DQds9PuX41a5BiuewGHDZeSAuzeIZFWu; rur=PRN; ig_pr=2; ig_vw=1235' \
    -H 'Connection: keep-alive' --data 'password=asdfasdfasdfasd&phone_number=%2B380971234588&username=bond776527&first_name=bond77&client_id=WOUeqAAEAAFrY7VM1GmjUKbEwxXL'

    curl 'https://www.instagram.com/accounts/send_signup_sms_code_ajax/' \
    -H 'Host: www.instagram.com' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:52.0) Gecko/20100101 Firefox/52.0' \
    -H 'Accept: */*' \
    -H 'Accept-Language: en-US,en;q=0.5' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'X-CSRFToken: DQds9PuX41a5BiuewGHDZeSAuzeIZFWu' \
    -H 'X-Instagram-AJAX: 1' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -H 'Referer: https://www.instagram.com/' \
    -H 'Cookie: mid=WOUeqAAEAAFrY7VM1GmjUKbEwxXL; csrftoken=DQds9PuX41a5BiuewGHDZeSAuzeIZFWu; rur=PRN; ig_pr=2; ig_vw=1235' \
    -H 'Connection: keep-alive' --data 'client_id=WOUeqAAEAAFrY7VM1GmjUKbEwxXL&phone_number=%2B380971234588'

}

login () {
    name=$1
    printf "\n L "
    curl 'https://www.instagram.com/accounts/login/ajax/' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en-US,en;q=0.5' \
    -H 'Accept: */*' \
    -H 'Connection: keep-alive' --data 'username='$name'&password=assdfdffghfy' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -H 'Cookie: csrftoken=z11zdfgVptur41IFtbmWVa2KpRqqkzl8; mid=WOSd5gAEAAFcrF6vwEo6N1UVZQB0; rur=PRN; ig_vw=1235; ig_pr=2' \
    -H 'Host: www.instagram.com' \
    -H 'Referer: https://www.instagram.com/' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:52.0) Gecko/20100101 Firefox/52.0' \
    -H 'X-CSRFToken: z11zdfgVptur41IFtbmWVa2KpRqqkzl8' \
    -H 'X-Instagram-AJAX: 1' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -c /tmp/cookie.tmp
    # {"authenticated": true, "status": "ok", "user": true}
}

follow () {
    printf "\n F "
    curl 'https://www.instagram.com/web/friendships/1457977472/follow/' -X POST -H 'Referer: https://www.instagram.com/k.vitalina/' \
    -H 'Accept-Encoding: gzip, deflate, br' \
    -H 'Accept-Language: en-US,en;q=0.5' \
    -H 'Accept: */*' \
    -H 'Connection: keep-alive' \
    -H 'Content-Length: 0' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -H 'Host: www.instagram.com' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:52.0) Gecko/20100101 Firefox/52.0' \
    -H 'X-CSRFToken: '`grep csrftoken /tmp/cookie.tmp | awk '{print $7}'` \
    -H 'X-Instagram-AJAX: 1' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -b /tmp/cookie.tmp
}

if [ $action = '1' ]; then
    for i in $(seq 1 1); do
        name=$(create)
        echo "# $i, username: $name"
        login "$name"
        follow
    done
fi

if [ $action = '2' ]; then
    declare -a arr=("fcvk1491376852")
    for name in "${arr[@]}"; do
        login "$name"
        follow
    done
fi
