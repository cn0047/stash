<?php

$path = dirname(__FILE__);

$objects = new RecursiveIteratorIterator(
    new RecursiveDirectoryIterator($path),
    RecursiveIteratorIterator::SELF_FIRST
);
foreach ($objects as $file => $object) {
    $basename = $object->getBasename();
    if ($basename == '.' or $basename == '..') {
        continue;
    }
    if ($object->isDir()) {
        continue;
    }
    if (date('Y-m-d') === date('Y-m-d', $object->getCTime())) {
        echo $file.PHP_EOL;
    }
}
