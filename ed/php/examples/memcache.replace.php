<?php

$memcache_obj = memcache_connect('localhost', 11211);
$memcache_obj->set("test_key", date('Y-m-d H:i:s'), false, 30);
var_export([
    $memcache_obj->get("test_key")
]);
$memcache_obj->replace("test_key", new \DateTime(), false, 30);
var_export([
    $memcache_obj->get("test_key")
]);
