<?php
/**
 * PHP Mcrypt.
 * @link http://runnable.com/VNVD6w47Ou8ax68e/php-mcrypt
 */

$t = 'Hello World';
$iv_size = mcrypt_get_iv_size(MCRYPT_RIJNDAEL_128, MCRYPT_MODE_CBC);
$iv = mcrypt_create_iv($iv_size, MCRYPT_RAND);
$c = mcrypt_encrypt(MCRYPT_RIJNDAEL_128, '123', $t, MCRYPT_MODE_CBC, $iv);
print($c . PHP_EOL);
$t = mcrypt_decrypt(MCRYPT_RIJNDAEL_128, '123', $c, MCRYPT_MODE_CBC, $iv);
print($t . PHP_EOL . PHP_EOL);

$iv_size = mcrypt_get_iv_size(MCRYPT_BLOWFISH, MCRYPT_MODE_ECB);
$iv = mcrypt_create_iv($iv_size, MCRYPT_DEV_URANDOM);
$c = mcrypt_encrypt(MCRYPT_BLOWFISH, '123', utf8_encode($t), MCRYPT_MODE_ECB, $iv);
print($c . PHP_EOL);
$t = mcrypt_decrypt(MCRYPT_BLOWFISH, '123', $c, MCRYPT_MODE_ECB, $iv);
print($t . PHP_EOL . PHP_EOL);
