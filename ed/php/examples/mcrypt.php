<?php
/**
 * PHP Mcrypt.
 *
 * @link http://runnable.com/VNVD6w47Ou8ax68e/php-mcrypt
 */

$t = 'Hello World';
$key = '1234567';

$ivSize = mcrypt_get_iv_size(MCRYPT_RIJNDAEL_128, MCRYPT_MODE_CBC);
$iv = mcrypt_create_iv($ivSize, MCRYPT_RAND);
$c = mcrypt_encrypt(MCRYPT_RIJNDAEL_128, $key, $t, MCRYPT_MODE_CBC, $iv);
print($c . PHP_EOL);
$t = mcrypt_decrypt(MCRYPT_RIJNDAEL_128, $key, $c, MCRYPT_MODE_CBC, $iv);
print($t . PHP_EOL . PHP_EOL);

/**
 * BAD EXAMPLE!!!
 *
 * MCRYPT_BLOWFISH is old and insecure!!! Use AES!!!
 */
$ivSize = mcrypt_get_iv_size(MCRYPT_BLOWFISH, MCRYPT_MODE_ECB);
$iv = mcrypt_create_iv($ivSize, MCRYPT_DEV_URANDOM);
$c = mcrypt_encrypt(MCRYPT_BLOWFISH, $key, utf8_encode($t), MCRYPT_MODE_ECB, $iv);
print($c . PHP_EOL);
$t = mcrypt_decrypt(MCRYPT_BLOWFISH, $key, $c, MCRYPT_MODE_ECB, $iv);
print($t . PHP_EOL . PHP_EOL);
