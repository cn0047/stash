#!/bin/sh

readonly URL='http://www.work.ua/jobs-kyiv-php/?days=123&page='

toOpen=''
i=0
while [ true ]; do
    i=$(($i+1))
    curl -s -o /tmp/jobSeeker.tmp "$URL$i"
    is404=$((`grep 'card-logotype' -ric /tmp/jobSeeker.tmp`))
    if [ $is404 -eq 0 ]; then
        /usr/bin/chromium-browser $toOpen
        exit
    fi
    for uri in `grep '\/jobs\/[0-9]+\/' -Erio /tmp/jobSeeker.tmp`; do
        curl -s -o /tmp/jobSeeker.item.tmp "http://www.work.ua$uri"
        isInteresting=$((`grep 'remote|удален|відд?ален' -Eric /tmp/jobSeeker.item.tmp`))
        exclude=$((`grep '1С Bitrix|на удаленную работу не рассматриваются' -Eric /tmp/jobSeeker.item.tmp`))
        if [ $isInteresting -gt 0 ]; then
            if [ $exclude -eq 0 ]; then
                toOpen="$toOpen http://www.work.ua$uri"
            fi
        fi
        rm /tmp/jobSeeker.item.tmp
        printf "\ri = $i, uri = $uri"
    done
    rm /tmp/jobSeeker.tmp
done
