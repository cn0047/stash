php
-

*PHP 5.1.6*

````php
<?php
$output = `ls -al`;
echo "<pre>$output</pre>";
````

tick - это событие, которое происходит для каждых N-инструкций нижнего уровня, выполненных синтаксическим анализатором в пределах блока declare.

````php
<?php
const foo = 200;
echo foo."\n";
echo constant('foo')."\n";
````

````php
<?php
// Colorizing php cli scripts
echo "\033[31mred\033[37m ";
echo "\033[32mgreen\033[37m ";
echo "\033[41;30mblack on red\033[40;37m ";
````