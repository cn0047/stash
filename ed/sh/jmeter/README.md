jmeter
-

[docs](http://jmeter.apache.org/usermanual/get-started.html)

jmeter written in java.

````sh
# random hash wiht length 10 chars
${__javaScript(Math.random().toString(36).toUpperCase().substring(16))}

# ENV variable
${__P('host')}
${__P('port', 3000)}
````

````sh
jmeter -n -t tests/jmeter/one.jmx \
  -Jjmeter.save.saveservice.output_format=xml \
  -Jjmeter.save.saveservice.samplerData=true \
  -Jjmeter.save.saveservice.url=false -Jjmeter.save.saveservice.requestHeaders=false \
  -Jjmeter.save.saveservice.response_data=true -Jjmeter.save.saveservice.responseHeaders=false \
  -l /tmp/jmeter.result -j /tmp/jmeter.log \
  -Jhost=api.h.dev \
````
