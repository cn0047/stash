<?php

$file = '/var/log/nginx/access.log';
$pos = 0;
$i = 0;
if (!is_readable($file)) {
    throw new \RuntimeException("Can't read file: {$file}");
}
$f = fopen($file, 'r');
$size = filesize($file);
if ($pos <= $size) {
    fseek($f, $pos);
}
while (($line = fgets($f)) !== false) {
    var_dump($line);
    $i++;
}
fclose($f);
echo "size: $size, i: $i".PHP_EOL;
