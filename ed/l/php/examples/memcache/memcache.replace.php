<?php

$mc = memcache_connect('xmemcached', 11211);
var_dump($mc->replace('simple-test', 3, false, 5));
