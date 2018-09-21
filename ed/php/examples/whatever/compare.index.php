<?php

header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept');

if ($_SERVER['QUERY_STRING'] === 'busy') {
    sleep(7);
}
echo '[php] It works!';

// curl 'http://localhost:80/index.php?busy'
