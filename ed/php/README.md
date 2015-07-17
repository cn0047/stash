PHP
-
<br>*PHP 5.5.9*
<br>*PHP 5.4.35*
<br>*PHP 5.3.3*
<br>*PHP 5.1.6*

https://dev.twitter.com/
|
http://www.phpdeveloper.org/
|
http://php.net/conferences/index.php
|
http://git.php.net/

/etc/php5/apache2/php.ini

`<![CDATA[]]>` - actual for XHTML.

php -S localhost:8000 index.php

php -d short_open_tag=1 x.php

````php
// Primitive validation
preg_match('/^\d{4}(-\d{2}){2}$/', $args['date']); // 2015-06-10
````

####Exceptions
````php
throw new ErrorException(404);

throw new BadFunctionCallException(404);
throw new BadMethodCallException(404);

throw new InvalidArgumentException(404);
throw new LengthException(404);
throw new OutOfBoundsException(404);
throw new OutOfRangeException(404);
throw new OverflowException(404);
throw new RangeException(404);
throw new UnexpectedValueException(404);

throw new DomainException(404);
throw new LogicException(404);
throw new RuntimeException(404);
throw new UnderflowException(404);
````

####phpCs
````
phpcs -sw --standard=PSR2 file.php

curl -OL https://squizlabs.github.io/PHP_CodeSniffer/phpcs.phar
chmod +x phpcs.phar
cp phpcs.phar /usr/local/bin/phpcs
cp phpcs.phar /usr/bin/phpcs
````

####[phpDoc](http://www.phpdoc.org/docs/latest/index.html)
````
phpdoc run -d . -t doc
````

````php
<?php
$output = `ls -al`;
echo "<pre>$output</pre>";
````

####xDebug
````
# Debug
# ?XDEBUG_SESSION_START=sublime.xdebug
# .htaccess
php_value xdebug.remote_enable 1
# php_value xdebug.remote_connect_back 1
# php_value xdebug.remote_host localhost
# php_value xdebug.remote_handler dbgp
# php_value xdebug.remote_port 9000
# php_value xdebug.remote_log '/tmp/xdebug.log'

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
````
vim ~/.bashrc

export PHP_IDE_CONFIG="serverName=trunk-007" 
export XDEBUG_CONFIG="idekey=PHPSTORM remote_host={192.168.13.58} remote_enable=1 remote_autostart=1" 
export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host={127.0.0.1} remote_enable=1 remote_autostart=1"

source ~/.bashrc
````

####debug
````php
echo '<script>console.log('.json_encode($_REQUEST).')</script>';

array_walk(debug_backtrace(), create_function('$v', '
    file_put_contents("/tmp/debug.tmp", sprintf("%s -> %s():%s\n", $v["file"], $v["function"], $v["line"]), FILE_APPEND); /// tail -f /tmp/debug.tmp
'));
foreach (debug_backtrace() as $v) {
    file_put_contents('/tmp/debug.tmp', $v['file'].' -> '.$v['function'].'():'.$v['line']."\n", FILE_APPEND); /// tail -f /tmp/debug.tmp
}

(mt_rand(0, 10) > 1) or var_dump(200);

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

$foo = '5bar';
settype($foo, 'integer'); // $foo is now 5 (integer)

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

####Data Types
* Scalar types:
    * Boolean.
    * Integer.
    * Float (floating-point number, aka double).
    * String.
* Compound types:
    * Array.
    * Object.
* Special types:
    * Resource.
    * NULL.
* Pseudo-types:
    * Mixed.
    * Number.
    * Callback (aka callable).

##History of PHP

http://php.net/manual/en/appendices.php

####PHP 5.6.x
* Constant scalar expressions.
* Variadic functions via ...
* Argument unpacking via ...
* Exponentiation via **
* Use function and use const.
* Phpdbg.
* Default character encoding.
* Files larger than 2 gigabytes in size are now accepted.

####PHP 5.5.x
* Generators.
* Try-catch blocks now support a finally.
* New password hashing API.
* Foreach now supports list().
* Empty() supports arbitrary expressions.
* Array and string literal dereferencing.
* Class name resolution via ::class. `ClassName::class`
* OPcache extension added.

####PHP 5.4.x
* Traits.
* Short array syntax.
* Function array dereferencing has been added. `foo()[0]`
* Closures now support `$this`.
* Class member access on instantiation has been added. `(new Foo)->bar()`
* Binary number format.

####PHP 5.3.x
* Namespaces.
* Late Static Bindings.
* Native Closures.
* Nowdoc syntax is now supported, similar to Heredoc syntax, but with single quotes.
* Constants can now be declared outside a class using the const keyword.

####PHP 5.2.x

####PHP 5.1.x
