Logstash
-

````sh
sudo service logstash configtest
````

## Config

Filters are applied in the order in which they appear in config.

````sh
# {"ua":{"os":"windows"}} in config it will be
[ua][os]
# for example
output {
  statsd {increment => "apache.%{[ua][os]}"}
}
````

Simple test:

````sh
/opt/logstash/bin/logstash -e 'input { stdin {} } output { stdout {} }'

````

Simple elasticsearch config:

````sh
input { file { path => "/tmp/logstash.txt" } } output { elasticsearch { hosts => ["localhost:9200"] } }
````
````sh
/opt/logstash/bin/logstash -f /etc/logstash/conf.d/my.conf
````

Nginx conf:

````sh
input { file { path => "/var/log/nginx/access.log" } } output { elasticsearch { hosts => ["localhost:9200"] index => "nginx" } }
````
