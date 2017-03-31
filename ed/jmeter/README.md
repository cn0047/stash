jmeter
-

````
jmeter -n -t tests/jmeter/one.jmx \
  -Jjmeter.save.saveservice.output_format=xml \
  -Jjmeter.save.saveservice.samplerData=true \
  -Jjmeter.save.saveservice.url=false -Jjmeter.save.saveservice.requestHeaders=false \
  -Jjmeter.save.saveservice.response_data=true -Jjmeter.save.saveservice.responseHeaders=false \
  -l /tmp/jmeter.result -j /tmp/jmeter.log \
  -Jhost=api.h.dev \
````

http://jmeter.apache.org/usermanual/get-started.html
