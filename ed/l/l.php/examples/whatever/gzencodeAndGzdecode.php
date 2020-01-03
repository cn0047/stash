<?php

$compressed = gzencode('Compress me', 9);
var_dump([
    $compressed,
    gzencode('sms_id=123'),
]);
$uncompressed = gzdecode($compressed);
var_dump([
    $uncompressed,
    gzinflate(substr($compressed, 10, -8)),
]);
