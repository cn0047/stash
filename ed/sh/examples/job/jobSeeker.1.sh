#!/bin/sh

readonly START_FROM_ID_OFFSET=22594

args="$@"
idFrom=$1

if [ -z "$idFrom" ]; then
    echo "Finding id"
    i=0
    while [ true ]; do
        id=$(($START_FROM_ID_OFFSET+$i))
        i=$(($i+1))
        curl -is -o /tmp/jobSeeker.tmp "http://jobs.site.com/vacancies/$id/"
        is404=$((`grep 'HTTP/1.1 404 NOT FOUND' -ric /tmp/jobSeeker.tmp`))
        printf "\ri = $i, id = $id, is 404 = $is404"
        if [ $is404 -eq 0 ]; then
            printf "\nFound id - $id\n"
            break
        fi
        rm /tmp/jobSeeker.tmp
        is404=0
    done
else
    echo "Start from $idFrom"
    i=0
    while [ true ]; do
        id=$(($idFrom+$i))
        i=$(($i+1))
        url="http://jobs.site.com/vacancies/$id/"
        curl -is -o /tmp/jobSeeker.tmp "$url"
        is404=$((`grep 'HTTP/1.1 404 NOT FOUND' -ric /tmp/jobSeeker.tmp`))
        isPhp=$((`grep 'php|пхп' -Eric /tmp/jobSeeker.tmp`))
        isRemote=$((`grep 'remote|удален|відд?ален' -Eric /tmp/jobSeeker.tmp`))
        printf "404 = $is404, php = $isPhp, remote = $isRemote, url = $url \n"
        if [ $is404 -eq 1 ]; then
            printf "Found 404.\n"
            break
        fi
        rm /tmp/jobSeeker.tmp
        is404=0
    done
fi
