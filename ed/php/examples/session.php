<?php

$ttl = 36000;
session_save_path('/tmp/mySessionPath');
ini_set('session.gc_maxlifetime', $ttl);
ini_set('session.cookie_lifetime', $ttl);
ini_set('session.cache_expire', $ttl);
ini_set('session.use_only_cookies', 1);
ini_set('session.gc_probability', 1);
ini_set('session.gc_probability', 1);
