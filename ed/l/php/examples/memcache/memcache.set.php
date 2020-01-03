<?php

$mc = memcache_connect('xmemcached', 11211);
// Won't update expiration time if key exists.
var_dump($mc->set('simple-test', 2, false, 5));
