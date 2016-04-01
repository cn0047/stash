kahlan
-

2.4

````
vendor/bin/kahlan --coverage="One"

vendor/bin/kahlan --pattern=spec/suite/services/v1/CaptionsSpec.php  --coverage="W3\Ziipr\Services\v1\Captions"
vendor/bin/kahlan --coverage="W3\Ziipr\Services\v1" --istanbul="coverage.json"

vendor/bin/kahlan --coverage=4 --istanbul="coverage.json"
istanbul report
/usr/bin/chromium-browser coverage/lcov-report/index.html
````
