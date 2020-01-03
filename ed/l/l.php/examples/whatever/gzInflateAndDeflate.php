<?php

$compressed = gzdeflate('Compress me', 9);
var_export([
    mb_strlen($compressed),
    $compressed,
    base64_encode($compressed),
]);
$uncompressed = gzinflate($compressed);
var_export([
    mb_strlen($uncompressed),
    $uncompressed,
    gzinflate(base64_decode('c87PLShKLS5WyE0FAA==')),
]);
