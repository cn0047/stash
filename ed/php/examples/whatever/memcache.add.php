<?php

$mc = memcache_connect('xmemcached', 11211);
// Return false if key exists.
var_dump($mc->add('simple-test', 1, false, 10));
