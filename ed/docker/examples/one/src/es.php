<?php

$response = `curl -s -XGET es:9200`;
$br = PHP_SAPI === 'cli' ? PHP_EOL : '<br>';
print("RESPONSE: $br $response");
