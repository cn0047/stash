#!/bin/sh

/usr/bin/chromium-browser "http://rabota.ua/ua/jobsearch/vacancy_list?regionId=1&keyWords=PHP&scheduleId=3&period=3&lastdate=24.03.2016"
exit

readonly DATE=`date "+%d.%m.%Y" -d "7 days ago"`
# Kiev
readonly URL="http://rabota.ua/jobsearch/vacancy_list?regionId=1&keyWords=php&period=3&lastdate=$DATE&pg=" # page
# All UA
# readonly URL="http://rabota.ua/jobsearch/vacancy_list?keyWords=php&period=3&lastdate=$DATE&pg=" # page

toOpen=''
i=0
while [ true ]; do
    i=$(($i+1))
    curl -s -o /tmp/jobSeeker.tmp "$URL$i"
    is404=$((`grep 'не нашли ничего' -Eric /tmp/jobSeeker.tmp`))
    if [ $is404 -gt 0 ]; then
        echo "\n"
        echo $toOpen
        # /usr/bin/chromium-browser $toOpen
        exit
    fi
    for uri in `grep '\/company[0-9]+\/vacancy[0-9]+' -Erio /tmp/jobSeeker.tmp`; do
        curl -s -o /tmp/jobSeeker.item.tmp "http://rabota.ua$uri"
        isInteresting=$((`grep 'remote|удален|відд?ален' -Eric /tmp/jobSeeker.item.tmp`))
        exclude=$((`grep '1С|удаленную работу не рассматрива|Вариант работы удаленно не рассматриваем|\/zapros\/php-программист-удаленно\/киев' -Eric /tmp/jobSeeker.item.tmp`))
        addedToHeap=0
        if [ $isInteresting -gt 0 ]; then
            if [ $exclude -eq 0 ]; then
                addedToHeap=1
                if [ "$toOpen" != *"$uri"* ]; then
                    toOpen="$toOpen http://rabota.ua$uri"
                fi
            fi
        fi
        printf "\rpage = $i, isInteresting = $isInteresting, exclude = $exclude, addedToHeap = $addedToHeap uri = $uri"
        rm /tmp/jobSeeker.item.tmp
    done
    rm /tmp/jobSeeker.tmp
done
