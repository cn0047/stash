<?php

$mc = memcache_connect('xmemcached', 11211);
// Return false if key NOT exists.
var_dump($mc->delete('simple-test'));
