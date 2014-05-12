php
-

*PHP 5.1.6*
####Flashback
````php
<?php

$_SERVER['HTTP_ACCEPT_LANGUAGE'];

$PHP_SELF;

mysql_real_escape_string();

get_defined_constants(true);

setcookie('TestCookie', $value);
$_COOKIE['TestCookie'];

header();

settype($var, $type);

gzuncompress(gzcompress($data));

filter_var();
filter_list();

highlight_string();

system();

set_include_path();

memory_get_usage();
memory_get_peak_usage();

is_a($object, $class_name);

get_class();
get_parent_class();
get_class_methods($class_name);
get_class_vars($class_name);
get_object_vars($object);
property_exists($class, $property);

tempnam($dir, $prefix);
````

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