Grafana
-

````sh
url=~.*needle.*
````

#### TraceQL

````sh
{ resource.service.name =~ ".+" && .rpc.method = "GetMyData" }
````
