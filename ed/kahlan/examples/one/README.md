One
-

````
vendor/bin/kahlan --istanbul="coverage.json"
istanbul report
/usr/bin/chromium-browser coverage/lcov-report/index.html

vendor/bin/kahlan --coverage='One\Text::get()'
vendor/bin/kahlan --coverage='One\Text'
vendor/bin/kahlan --coverage='One'
````
