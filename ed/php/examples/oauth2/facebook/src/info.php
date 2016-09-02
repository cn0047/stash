<pre>
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

echo '<h3>Metadata</h3>';
try {
    // Returns a `Facebook\FacebookResponse` object
    $response = $fb->get(
        '/me?fields=id,name,first_name,middle_name,last_name,email,locale,location,hometown,gender,picture,birthday',
        $_GET['token']
    );
} catch(Facebook\Exceptions\FacebookResponseException $e) {
    echo 'Graph returned an error: ' . $e->getMessage();
    exit;
} catch(Facebook\Exceptions\FacebookSDKException $e) {
    echo 'Facebook SDK returned an error: ' . $e->getMessage();
    exit;
}

$user = $response->getGraphUser();
var_export($user);
var_export($user->getLocation());
var_export($user->getPicture());
