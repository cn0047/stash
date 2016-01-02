#!/bin/sh

# readonly URL='http://www.work.ua/jobs-kyiv-php/?days=123&page='
readonly URL='http://www.work.ua/jobs-php/?days=123&page='

toOpen=''
i=0
while [ true ]; do
    i=$(($i+1))
    curl -s -o /tmp/jobSeeker.tmp "$URL$i"
    is404=$((`grep 'card-logotype' -ric /tmp/jobSeeker.tmp`))
    if [ $is404 -eq 0 ]; then
        echo $toOpen
        /usr/bin/chromium-browser $toOpen
        exit
    fi
    for uri in `grep '\/jobs\/[0-9]+\/' -Erio /tmp/jobSeeker.tmp`; do
        curl -s -o /tmp/jobSeeker.item.tmp "http://www.work.ua$uri"
        isInteresting=$((`grep 'remote|удален|відд?ален' -Eric /tmp/jobSeeker.item.tmp`))
        exclude=$((`grep '1С|удаленную работу не рассматрива|Вариант работы удаленно не рассматриваем' -Eric /tmp/jobSeeker.item.tmp`))
        addedToHeap=0
        if [ $isInteresting -gt 0 ]; then
            if [ $exclude -eq 0 ]; then
                addedToHeap=1
                if [ "$toOpen" != *"$uri"* ]; then
                    toOpen="$toOpen http://www.work.ua$uri"
                fi
            fi
        fi
        printf "\rpage = $i, isInteresting = $isInteresting, exclude = $exclude, addedToHeap = $addedToHeap uri = $uri"
        rm /tmp/jobSeeker.item.tmp
    done
    rm /tmp/jobSeeker.tmp
done
