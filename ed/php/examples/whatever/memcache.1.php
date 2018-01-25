<?php

$memcache_obj = memcache_connect('xmemcached', 11211);

/* procedural API */
memcache_add($memcache_obj, 'var_key_1', 'test variable: '.uniqid(), false, 30);
/* OO API */
$memcache_obj->add('var_key_2', 'test variable: '.uniqid(), false, 30);

var_export([
    /* procedural API */
    memcache_get($memcache_obj, 'var_key_1'),
    /* OO API */
    $memcache_obj->get('var_key_2'),
]);
