<?php

$redis = new Redis();
$redis->connect('127.0.0.1', 6379);

var_export($redis->sMembers('srelationships:user10:following'));
var_export($redis->sMembers('srelationships:user10:followers'));

var_export($redis->hGetAll('hrelationships:user10:following'));
var_export($redis->hGetAll('hrelationships:user10:followers'));
