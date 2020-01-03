<?php

use Persistence\User;

require_once __DIR__ . '/vendor/autoload.php';

$user = (new User())->getById(1);
var_export($user);
