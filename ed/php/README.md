php
-
*PHP 5.1.6*

url
````
    ? %3F
    & %26
    = %3D
    : %3A
    / %2F
    - -
    _ _
````

`<![CDATA[]]>`

**phpDoc** `phpdoc run -d . -t doc`

**phpUnit** `/usr/share/php/PHPUnit/Extensions/SeleniumTestCase.php`

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

####xDebug
````
# Debug
# ?XDEBUG_SESSION_START=sublime.xdebug
# .htaccess
php_value xdebug.remote_enable 1
php_value xdebug.remote_connect_back 1
php_value xdebug.remote_host localhost
php_value xdebug.remote_handler dbgp
php_value xdebug.remote_port 9000

# Trace
xdebug_start_trace('~/debug/trace.xt');
xdebug_stop_trace();

# Profiling
# .htaccess
php_value xdebug.profiler_enable 1
php_value xdebug.profiler_output_name xdebug.out.%t
php_value xdebug.profiler_output_dir /home/debug/d
php_value xdebug.profiler_enable_trigger 1
# Generated file open with KCachegrind.
````

####debug
````php
$t =microtime();
var_export(sprintf('%f', microtime()-$t));

error_reporting(E_ALL); // error_reporting(E_ALL & ~E_NOTICE);
ini_set('display_errors', 1);
ini_set('display_startup_errors','On');

// debug tables:
array_walk($d, create_function('&$i, $k, $c', 'if (empty($c)) {$c=array_keys($i);} $i="<tr><td>".implode("</td><td>",$i)."</td></tr>";'), &$c);
echo "<table border=1 cellspacing=0 cellpadding=3 bordercolor='#BFBFBF'><tr bgcolor='#ADD8E6' align=center><td>".implode("</td><td>",$c)."</td></tr>".implode("",$d)."</table>";
require_once SITE_PATH.'Table.php';
$tbl = new Console_Table();
$tbl->setHeaders(array_keys($d[0]));
foreach($d as $v){$tbl->addRow($v);}
echo '<pre>'.$tbl->getTable().'</pre>';

set_error_handler(create_function('$n, $s, $f, $l', 'var_export(array($n, $s, $f, $l));'));
set_error_handler(create_function('$n, $s, $f, $l', 'print("\033[01;31m ".$s." \033[0m \n");'));
set_error_handler(function ($no, $str, $file, $line) {
  mail('mail@mail.com', 'DBG|ERROR', var_export([$no, $str, $file, $line], 1));
});

php_sapi_name() == 'cli' ? print("\n$error\n") : pr($error);
````

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
