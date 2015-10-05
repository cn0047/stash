fail=0;
for i in $(cat /etc/supervisor/conf.d/monitor.conf | grep 'command=\/usr\/bin\/php \/var\/www\/.*\.php' -oE | grep '\/var\/www\/.*\.php' -oE);
do
    result=$((`ps aux | grep -v grep | grep $i -c`))
    if [ $result -gt 0 ]
    then
        fail=$((fail+0));
    else
        fail=$((fail+1));
    fi
done
if [ $fail -gt 0 ];
then
    echo 0;
else
    echo 1;
fi
