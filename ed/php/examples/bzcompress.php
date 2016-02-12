<?php

$compressed = bzcompress('Compress me');
var_export([
    $compressed,
]);
$uncompressed = bzdecompress('$compressed');
var_export([
    $uncompressed,
]);
