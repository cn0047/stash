<?php

error_reporting(E_ALL); // error_reporting(E_ALL & ~E_NOTICE);
ini_set('display_errors', 1);
ini_set('display_startup_errors','On');

require_once __DIR__ . '/../vendor/autoload.php';
require_once __DIR__ . '/config.php';

session_start();
date_default_timezone_set('America/Los_Angeles');

$fb = new Facebook\Facebook([
    'app_id' => APP_ID,
    'app_secret' => APP_SECRET,
    'default_graph_version' => 'v2.2',
]);

$helper = $fb->getRedirectLoginHelper();

$permissions = ['email'];
$loginUrl = $helper->getLoginUrl('http://' . HOST .'/fb-callback.php', $permissions);

echo '<a href="' . htmlspecialchars($loginUrl) . '">Log in with Facebook!</a>';
