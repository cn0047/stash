phpUnit
-
*PHPUnit 3.7*

/usr/share/php/PHPUnit
/usr/share/php/PHPUnit/Extensions/SeleniumTestCase.php

Stub - an object that provides predefined answers to method calls.
<br>Mock - an object on which you set expectations.

`
phpunit -d memory_limit=16M -c include/tests/conf.xml /tests/unit/
`

In each test should be just one assert.
