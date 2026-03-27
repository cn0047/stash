Grafana
-

#### TraceQL

````sh
{ resource.service.name =~ ".+" && .rpc.method = "GetMyData" }
````
