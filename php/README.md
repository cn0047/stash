php
-

*PHP 5.1.6*

````php
<?php
$output = `ls -al`;
echo "<pre>$output</pre>";
?>
````


tick - это событие, которое происходит для каждых N-инструкций нижнего уровня, выполненных синтаксическим анализатором в пределах блока declare.


````php
<?php
const foo = 200;
echo foo."\n";
echo constant('foo')."\n";
````
