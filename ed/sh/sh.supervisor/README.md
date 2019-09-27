Supervisor
-

#### Configuration file
````sh
/etc/supervisor/supervisord.conf
````

````sh
http://localhost:9001/

service supervisor status
service supervisord status
/etc/init.d/supervisor restart
/etc/init.d/supervisord.sh restart

sudo /usr/local/bin/supervisorctl status
sudo /usr/local/bin/supervisorctl restart all
sudo /usr/local/bin/supervisorctl reload # reload config
supervisorctl status
supervisorctl reload
supervisorctl restart all
supervisorctl stop all
````
