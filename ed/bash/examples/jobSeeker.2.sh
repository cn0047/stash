#!/bin/sh

readonly URL_1='http://jobs.com/vacancies/?city=x'
readonly URL_2='http://jobs.com/vacancies/xhr-load/?city=x&count_BAD=9999'

curl -s -o /tmp/jobSeeker.out.1.tmp -c /tmp/jobSeeker.cookie.tmp $URL_1
token=`grep 'csrftoken(.*)' -Erio /tmp/jobSeeker.cookie.tmp | awk '{print $2}'`
curl -s -X POST -o /tmp/jobSeeker.out.xhr.2.tmp -c /tmp/jobSeeker.cookie.tmp -b /tmp/jobSeeker.cookie.tmp -d count=100 -d csrfmiddlewaretoken=$token $URL_2
php -r 'var_export(json_decode(`cat /tmp/jobSeeker.out.xhr.2.tmp`, 1));' > /tmp/jobSeeker.out.2.tmp
rm `ls /tmp/jobSeeker.out.xhr.*.tmp`
subl `ls /tmp/jobSeeker.out*`

grep 'https?:\/\/.*\/vacancies\/[0-9]+' -Erin /tmp/jobSeeker.out.* | wc -l
grep 'https?:\/\/.*\/vacancies\/[0-9]+' -Erin /tmp/jobSeeker.out.*
grep 'jobs.com\/companies' -Erio /tmp/jobSeeker.out.*|wc -l

rm `ls /tmp/jobSeeker.*`
ls /tmp/jobSeeker.*

# $RANDOM
