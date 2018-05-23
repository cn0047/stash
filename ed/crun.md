crun
-

`crontab -e`
````
cat /etc/cron.d/
````

````
SHELL=/bin/bash
PATH=/bin:/usr/bin:/usr/local/bin
CRUN_REMOTE_HOST=Host
CRUN_EMAIL=mail@com.com
CRUN_WORK_DIR=/var/www/vhosts/host/htdocs

0 * * * * /bin/echo `date` >> /tmp/d.tmp

# crun lock file. Should be removed when execution fails...
rm /tmp/crun_user_at_server_or_host_3a30db060f74d9390a2eb6f8a92eab8d

grep cron /var/log/syslog
````

````
# restart on centos
sudo service crond restart
````
