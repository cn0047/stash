<?php

$date = '2015-07-15 07:10:26';
$date0 = new \DateTime($date);
$date1 = new \DateTime($date, new \DateTimeZone('Europe/Kiev'));
$date1->setTimezone(new \DateTimeZone('GMT'));
$date2 = new \DateTime($date, new \DateTimeZone('UTC'));
$date3 = new \DateTime($date, new \DateTimeZone('Europe/Kiev'));
$date3->setTimezone(new \DateTimeZone('UTC'));

var_export([
    $date,
    $date0->format('Y-m-d H:i:s'),
    $date1->format('Y-m-d H:i:s').' *** '.$date1->format('Y-m-d H:i:s eO T'),
    $date2->format('Y-m-d H:i:s'),
    $date3->format('Y-m-d H:i:s').' *** '.$date3->format('Y-m-d H:i:s eO T'),
]);

/*
array (
  0 => '2015-07-15 07:10:26',
  1 => '2015-07-15 07:10:26',
  2 => '2015-07-15 04:10:26 *** 2015-07-15 04:10:26 GMT+0000 GMT',
  3 => '2015-07-15 07:10:26',
  4 => '2015-07-15 04:10:26 *** 2015-07-15 04:10:26 UTC+0000 UTC',
)
*/
