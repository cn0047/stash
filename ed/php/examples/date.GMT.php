<?php

$date = '2015-12-15 07:00:26';
$dateCreated = new \DateTime($date, new \DateTimeZone('Europe/Kiev'));
$dateCreated->setTimezone(new \DateTimeZone('GMT'));

var_export([
    $date,
    $dateCreated->format('Y-m-d H:i:s'),
]);

/*
array (
  0 => '2015-12-15 07:00:26',
  1 => '2015-12-15 05:00:26',
)
*/
