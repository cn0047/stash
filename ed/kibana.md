Kibana
-

````
sudo service kibana status

curl http://localhost:5601
````

Search:
````
*
@log_group:nginx_prod and @log_stream:error_log
````
