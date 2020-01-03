<?php

require_once __DIR__ . '/vendor/autoload.php';

$eu = new VISA\EuropeanUnion();
$result = $eu->approve();
var_export($result);
