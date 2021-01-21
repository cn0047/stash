PHP
-
<br>7.1.7
<br>5.6.23
<br>5.5.9
<br>5.4.35
<br>5.3.3
<br>5.1.6
Since 1994.

PHP - is high-level (not assembler), dynamic (behavior determines in runtime)
weakly typed, object-oriented
and interpreted server-side scripting language.

PEAR - PHP Extension and Application Repository.
PECL - Php extension community library.
SPL - Standard PHP Library.

[doc](http://php.net/)
[developer](http://www.phpdeveloper.org)
[git](http://git.php.net)
[functions](http://php.net/manual/en/funcref.php)
[conferences](http://php.net/conferences/index.php)
[php-src](https://github.com/php/php-src)

````
declare(strict_types = 1);

error_reporting(E_ALL); // error_reporting(E_ALL & ~E_NOTICE);
ini_set('display_errors', 1);
ini_set('display_startup_errors','On');
````

````
/etc/php5/apache2/php.ini
/etc/php5/cli/php.ini
/etc/php5/fpm/php.ini
/usr/local/etc/php/7.1/php.ini

slowlog = /tmp/php.log
request_slowlog_timeout = 1
````

`<![CDATA[]]>` - actual for XHTML.

````sh
php -S localhost:8000 index.php
php -S localhost:8000 -t ./public
# on vagrant
php -S 0.0.0.0:80 index.php

php -d short_open_tag=1 x.php
php -d extension=xhprof.so foo.php
````

````php
print(true ? 'Yes' : 'No'); // Output: Yes
print('My Value' ?: 'No Value'); // PHP 5.3 (Equals to empty($var) ? 'default' : $var) Output: My Value
print(null ?? 'Default Value'); // PHP 7 (Equals to isset($var) ? $var : 'default') Output: Default Value
````

````php
// Primitive validation
ctype_digit($testcase) // string with integer value
preg_match('/^\d{4}(-\d{2}){2}$/', $args['date']); // 2015-06-10
(!filter_var($args['email'], FILTER_VALIDATE_EMAIL))
(!filter_var($args['url'], FILTER_VALIDATE_URL))

$error = json_last_error();
if ($error !== JSON_ERROR_NONE) {
    throw new RuntimeException("JSON decode error: $error");
}
````

Streams:
````
STDIN  'php://stdin'
STDOUT 'php://stdout'
STDERR 'php://stderr'

php://input # read-only
php://output # write-only
php://fd # direct access to the given file descriptor
php://memory
php://temp
php://filter
````

#### opCache

https://gist.github.com/ck-on/4959032/?ocp.php
https://github.com/rlerdorf/opcache-status
https://github.com/PeeHaa/OpCacheGUI
https://github.com/amnuts/opcache-gui :+1:

#### Closures

Anonymous functions - functions which have no specified name.
<br>They are most useful as the value of callback parameters.

Lambda function is an anonymous PHP function that can be stored in a variable
and passed as an argument to other functions or methods.

PHP automatically converts such expressions into instances of the Closure internal class.
<br>A closure is a lambda function that is aware of its surrounding context.

`this` in closure - context of function declaration. In global scope equals to NULL, in class - class itself.

#### Escaping from HTML

````php
<?= 'print this string' ?>

<?php if ($expression == true): ?>
    This will show if the expression is true.
<?php else: ?>
    Otherwise this will show.
<?php endif; ?>

<?php foreach([1, 2, 3] as $number): ?>
    <?=$number?>
<?php endforeach; ?>
````

#### Exceptions

````php
throw new ErrorException('Unrecoverable Error', 456);

throw new \BadFunctionCallException();
throw new \BadMethodCallException('Method Not Allowed', 405);

throw new \InvalidArgumentException('Not acceptable.', 406);
throw new \LengthException('Request Entity Too Large', 413);
throw new \OutOfBoundsException(400);
throw new \OutOfRangeException(400);
throw new \OverflowException(400);
throw new \RangeException(400);
throw new \UnexpectedValueException('Precondition Failed', 412);

throw new \DomainException('Failed Dependency', 424);
throw new \LogicException(404);
throw new \RuntimeException('Expectation Failed', 417);
throw new \UnderflowException(404);
````

#### phpCs

````sh
phpcs -sw --standard=PSR2 file.php
phpcs -sw --standard=PSR2 --colors dir

curl -OL https://squizlabs.github.io/PHP_CodeSniffer/phpcs.phar
chmod +x phpcs.phar
sudo cp phpcs.phar /usr/local/bin/phpcs
sudo cp phpcs.phar /usr/bin/phpcs
````

#### [phpDoc](http://www.phpdoc.org/docs/latest/index.html)

````sh
phpdoc run -d . -t doc
````

````php
<?php
$output = `ls -al`;
echo "<pre>$output</pre>";
````

#### xDebug

````
xdebug_debug_zval('a');
````

````
[xdebug]
;
zend_extension=/usr/lib/php5/20131226/xdebug.so
xdebug.remote_port=9000
xdebug.remote_enable=On
xdebug.remote_connect_back=On
xdebug.remote_log=/tmp/xdebug.log
xdebug.idekey=PHPSTORM
;
xdebug.profiler_enable=On
xdebug.profiler_output_nam=xdebug.profiler.%p
xdebug.profiler_output_dir=/var/www/html/public/rest-api/prof/

# Debug
# ?XDEBUG_SESSION_START=sublime.xdebug
# ?XDEBUG_SESSION_START=PHPSTORM
# .htaccess
php_value xdebug.remote_enable 1
# php_value xdebug.remote_connect_back 1
# php_value xdebug.remote_host localhost
# php_value xdebug.remote_handler dbgp
# php_value xdebug.remote_port 9000
# php_value xdebug.remote_log '/tmp/xdebug.log'

# Trace
xdebug_start_trace('/tmp/xdebug.trace'); // output will be in file /tmp/xdebug.trace.xt
xdebug_stop_trace();

# Profiling
# .htaccess
php_value xdebug.profiler_enable 1
php_value xdebug.profiler_output_name xdebug.profiler.%t
php_value xdebug.profiler_output_dir /home/debug/d
# Generated file open with KCachegrind or qcachegrind on OSX.
````
````sh
vim ~/.bashrc

export PHP_IDE_CONFIG="serverName=trunk-007" 
export XDEBUG_CONFIG="idekey=PHPSTORM remote_enable=1 remote_autostart=1"
export XDEBUG_CONFIG="idekey=PHPSTORM remote_host={192.168.56.101} remote_enable=1 remote_autostart=1"
export XDEBUG_CONFIG="idekey=sublime.xdebug remote_host={127.0.0.1} remote_enable=1 remote_autostart=1"

source ~/.bashrc
````

#### debug

````php
php -r 'var_export(json_decode(`curl http://country.io/iso3.json`, 1));'

file_put_contents('/tmp/debug.log', var_export(this, 1)."\n", FILE_APPEND); /// tail -f /tmp/debug.log

echo "\033[31m ".var_export($e->getMessage(), 1)." \033[0m\n";

echo '<script>console.log('.json_encode($_REQUEST).')</script>';

array_walk(debug_backtrace(), create_function('$v', '
    file_put_contents("/tmp/debug.log", sprintf("%s -> %s():%s\n", $v["file"], $v["function"], $v["line"]), FILE_APPEND); /// tail -f /tmp/debug.log
'));
foreach (debug_backtrace() as $v) {
    file_put_contents('/tmp/debug.log', $v['file'].' -> '.$v['function'].'():'.$v['line']."\n", FILE_APPEND); /// tail -f /tmp/debug.log
}

(mt_rand(0, 10) > 1) or var_dump(200);

$t = microtime();
var_export(sprintf('%f', microtime()-$t));

// debug tables:
array_walk($d, create_function('&$i, $k, $c', 'if (empty($c)) {$c=array_keys($i);} $i="<tr><td>".implode("</td><td>",$i)."</td></tr>";'), &$c);
echo "<table border=1 cellspacing=0 cellpadding=3 bordercolor='#BFBFBF'><tr bgcolor='#ADD8E6' align=center><td>".implode("</td><td>",$c)."</td></tr>".implode("",$d)."</table>";
require_once SITE_PATH.'Table.php';
$tbl = new Console_Table();
$tbl->setHeaders(array_keys($d[0]));
foreach($d as $v){$tbl->addRow($v);}
echo '<pre>'.$tbl->getTable().'</pre>';

set_error_handler(function ($code, $description) {
    throw new ErrorException($description, $code);
});
set_error_handler(create_function('$n, $s, $f, $l', 'var_export(array($n, $s, $f, $l));'));
set_error_handler(create_function('$n, $s, $f, $l', 'print("\033[01;31m ".$s." \033[0m \n");'));
set_error_handler(function ($no, $str, $file, $line) {
  mail('mail@mail.com', 'DBG|ERROR', var_export([$no, $str, $file, $line], 1));
});

php_sapi_name() == 'cli' ? print("\n$error\n") : pr($error);
````

#### Flashback

````php
<?php

$foo = '5bar';
settype($foo, 'integer'); // $foo is now 5 (integer)

$_SERVER['HTTP_ACCEPT_LANGUAGE'];

$PHP_SELF;

// Escape
$html['username'] = htmlentities($clean['username'], ENT_QUOTES, 'UTF-8');
htmlspecialchars();
mysql_real_escape_string();
addslashes();
escapeshellcmd();
escapeshellarg();

get_defined_constants(true);

setcookie('TestCookie', $value);
$_COOKIE['TestCookie'];

header();

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

#### Interface vs abstract class

Interface:
describes class behaviour at all,
implements, multiple implements,
can contains only declatations of public methods,
contains constants that cannot be overrided in child class.

Abstract class:
describes particular class,
extends,
can contains properties, constants and methods, must contains abstract methods protected or public,
cannot be instantiated.

#### Predefined variables

* $_GET
* $_POST
* $_FILES
* $_COOKIE
* $_REQUEST
* $_SESSION
* $_SERVER
* $_ENV
* $GLOBALS

#### Magic methods

* __construct()
* __destruct() - uset, unload from function, shutdown
* __call()
* __callStatic()
* __get()
* __set()
* __isset() - calling isset() or empty() on properties.
* __unset() - unset() properties.
* __sleep() - serialize()
* __wakeup()
* __toString()
* __invoke()
* __set_state() - var_export()
* __clone()
* __debugInfo()

#### Data Types

* Scalar types:
    * Boolean.
    * Integer.
    * Float (floating-point number, aka double).
    * String.
* Compound types:
    * Array.
    * Object (passes by ref).
* Special types:
    * Resource (passes by ref).
    * NULL.
* Pseudo-types:
    * Mixed.
    * Number.
    * Callback (aka callable).

Type declarations (aka type hints):

    * int
    * float
    * bool
    * string
    * array
    * iterable
    * callable
    * self (`instanceof` the given class)
    * object (php 7.2)

It's ok to write `float[]` in phpDoc.
