<?php

preg_match(
    '@^(?:http://)?([^/]+)@i',
    'http://www.php.net/index.html',
    $matches
);
var_export($matches);
/*
array (
  0 => 'http://www.php.net',
  1 => 'www.php.net',
)
*/
