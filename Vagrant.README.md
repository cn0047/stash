sudo vim /etc/hosts
127.0.0.1 laravel.one.dev


http://laravel.one.dev:8080/


# result=$(grep laravel.one.dev -rc /etc/hosts)
# if [ $result -gt 0 ]
# then
#     sudo echo -e "127.0.0.1 laravel.one.dev\n`cat /etc/hosts`" > /etc/hosts
# fi
