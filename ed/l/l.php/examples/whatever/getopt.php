<?php

// php ed/php/examples/getopt.php  --sId=8 --dateFrom='2011-01-01' --dateTo='2011-02-01'
$options = getopt('', array('sId:', 'dateFrom:', 'dateTo:'));
var_export($options);

/*
array (
  'sId' => '8',
  'dateFrom' => '2011-01-01',
  'dateTo' => '2011-02-01',
)
*/
