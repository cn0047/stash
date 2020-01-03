<?php

require_once './vendor/autoload.php';

$clientId = '';
$clientSecret = '';
$redirectUrl = '';
$homeUrl = '';
$applicationName = '';

session_start();

$googleClient = new Google_Client();
$googleClient->setApplicationName($applicationName);
$googleClient->setClientId($clientId);
$googleClient->setClientSecret($clientSecret);
$googleClient->setRedirectUri($redirectUrl);
$googleClient->addScope('https://www.googleapis.com/auth/userinfo.profile');
$googleClient->addScope('https://www.googleapis.com/auth/userinfo.email');

if (isset($_GET['code'])) {
    $googleClient->authenticate($_GET['code']);
    $_SESSION['token'] = $googleClient->getAccessToken();
    header("Location: $redirect_url");
}
if (isset($_SESSION['token'])) {
    $googleClient->setAccessToken($_SESSION['token']);
}
if ($googleClient->getAccessToken()) {
    $plus = new Google_Service_Plus($googleClient);
    $userProfile = $plus->people->get('me');
    echo '<pre>';
    var_export($userProfile);
    echo '</pre>';
} else {
    $authUrl = $googleClient->createAuthUrl();
}
if (isset($authUrl)) {
    echo "<a href='$authUrl'>login</a>";
}
