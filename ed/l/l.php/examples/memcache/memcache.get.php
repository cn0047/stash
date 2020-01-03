<?php

$mc = memcache_connect('xmemcached', 11211);
while (true) {
    printf("at: %f value: %s\n", microtime(true), var_export($mc->get('simple-test'), true));
    sleep(1);
}
