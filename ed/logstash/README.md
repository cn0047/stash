Logstash
-

````
sudo service logstash configtest
````

## Config

Filters are applied in the order in which they appear in config.

````
# {"ua":{"os":"windows"}} in config it will be
[ua][os]
# for example
output {
  statsd {increment => "apache.%{[ua][os]}"}
}
````
````
input { file { path => "/tmp/logstash.txt" } } output { elasticsearch { host => "localhost" protocol => "http" } }
````
````
input {
  file {
    type => "nginx-access"
    path => "/var/log/nginx/access.log"
  }
}
filter {
}
output {
  elasticsearch {
    index => "nginx"
    type => "accessLog"
    embedded => false
    host => "localhost"
    protocol => "http"
    port => "9200"
  }
}

````
````
# siple test
/opt/logstash/bin/logstash -e 'input { stdin { } } output { stdout {} }'

/opt/logstash/bin/logstash -f /etc/logstash/conf.d/my.conf
````
