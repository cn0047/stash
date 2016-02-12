<?php

$time = mktime(0, 0, 0, 1, 1, 1998);
echo date_default_timezone_get()."\n";
echo   date("M d Y H:i:s", $time)."\n";
echo gmdate("M d Y H:i:s", $time)."\n";

/*
Europe/Helsinki
Jan 01 1998 00:00:00
Dec 31 1997 22:00:00
*/

echo "\n";

date_default_timezone_set('UTC');
$time = mktime(0, 0, 0, 1, 1, 1998);
echo date('M d Y H:i:s', $time)."\n";
date_default_timezone_set('Europe/Helsinki');
echo date('M d Y H:i:s', $time)."\n";

/*
Jan 01 1998 00:00:00
Jan 01 1998 02:00:00
*/
