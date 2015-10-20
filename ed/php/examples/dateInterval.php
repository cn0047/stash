<?php

$begin = new DateTime('00:00:00');
$end = new DateTime('23:59:59');
$interval = new DateInterval('PT2H');
$daterange = new DatePeriod($begin, $interval ,$end);
foreach($daterange as $date){
    echo $date->format('H:i:s').PHP_EOL;
}
/*
00:00:00
02:00:00
04:00:00
06:00:00
08:00:00
10:00:00
12:00:00
14:00:00
16:00:00
18:00:00
20:00:00
22:00:00
*/
