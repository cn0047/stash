<pre>
<?php

$_GET['token'] = 'EAAY4jeoZCK0sBAIJIERu4bEa8k4Em1PivZCU08PuTsfSzX4ZA1FDJcTxieqKlTOF2jrZAyjHlTpsDZBJhJhNEOZBbK5Awp8tKdeJJt837ka1TCZB88EcxoMtj8tcC5cFTk3UKsRmXWsOEv9x2FCNWv5FkGeijWb7o8ZD';
// EAAChFiY6eN4BAK3O6NAmnmA4VrANAQ34XatlInxvZCe9UVvYyZCqoPPE8bAuhiKFtGdpkWp2TBXxtSqGyxZB2qZBK9XxMpbCMqQ3fmqCnQ3oZChDystymeB1BIZC2CEmaxC1QHgu5R1gkQEJXGiQ1NiHdlbMKbOSWitwn9J7uvsEZA1pHEERlEc7Vy9cNtZAZCYfTty3859bzjKQCJTF2WRM6bZBrWUZCnLvweOqzAMHlGcZAQZDZD

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
        '/me?fields=id,name,first_name,middle_name,last_name,email,locale,location,hometown,gender,picture.width(800).height(800),birthday',
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
