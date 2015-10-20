<?php

$compressed = gzdeflate('Compress me', 9);
var_dump($compressed);
echo mb_strlen($compressed).PHP_EOL;

$uncompressed = gzinflate($compressed);
var_dump($uncompressed);
echo mb_strlen($uncompressed).PHP_EOL;
