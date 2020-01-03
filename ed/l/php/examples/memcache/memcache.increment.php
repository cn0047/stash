<?php

$mc = memcache_connect('xmemcached', 11211);
var_dump($mc->increment('simple-test', 1));
