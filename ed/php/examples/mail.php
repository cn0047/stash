<?php

mail('cnfxlr@gmail.com', 'TEST 0', 'TEST');

$from = 'cn2010_1@yahoo.com';
$headers = "From: $from\r\n";
$headers .= "Reply-To: cn2010_1@yahoo.com\r\n";
mail('cnfxlr@gmail.com', 'TEST 2', 'TEST', $headers, "-f$from");

mail('cnfxlr@gmail.com', 'TEST 3', 'TEST', "Reply-To: cn2010_1@yahoo.com\r\n");
