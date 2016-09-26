Logstash
-

````
sudo service logstash configtest
````
````
input {
  file {
    type => "nginx-access"
    path => "/var/log/nginx/access.log"
  }
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
/opt/logstash/bin/logstash -e 'input { stdin { } } output { stdout {} }'
````
