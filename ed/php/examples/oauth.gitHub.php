<?php
define('OAUTH2_CLIENT_ID', '');
define('OAUTH2_CLIENT_SECRET', '');

$authorizeURL = 'https://github.com/login/oauth/authorize';
$tokenURL = 'https://github.com/login/oauth/access_token';
$apiURLBase = 'https://api.github.com/';

session_start();

// Start the login process by sending the user to Github's authorization page
if($_GET['action'] == 'login') {
    // Generate a random hash and store in the session for security
    $_SESSION['state'] = hash('sha256', microtime(TRUE).rand().$_SERVER['REMOTE_ADDR']);
    unset($_SESSION['access_token']);
    $params = array(
        'client_id' => OAUTH2_CLIENT_ID,
        'redirect_uri' => 'http://kiev-pug-dev.890m.com/l.php',
        'scope' => 'user',
        'state' => $_SESSION['state']
    );
    // Redirect the user to Github's authorization page
    header('Location: ' . $authorizeURL . '?' . http_build_query($params));
    die();
}

// When Github redirects the user back here, there will be a "code" and "state" parameter in the query string
if($_GET['code']) {
    // Verify the state matches our stored state
    if(!$_GET['state'] || $_SESSION['state'] != $_GET['state']) {
        header('Location: ' . $_SERVER['PHP_SELF']);
        die();
    }
    // Exchange the auth code for a token
    $token = apiRequest($tokenURL, array(
        'client_id' => OAUTH2_CLIENT_ID,
        'client_secret' => OAUTH2_CLIENT_SECRET,
        'redirect_uri' => 'http://kiev-pug-dev.890m.com/l.php',
        'state' => $_SESSION['state'],
        'code' => $_GET['code']
    ));
    $_SESSION['access_token'] = $token->access_token;
    header('Location: ' . $_SERVER['PHP_SELF']);
}

if($_SESSION['access_token']) {
    $user = apiRequest($apiURLBase . 'user');
    echo '<h3>Logged In</h3>';
    echo '<h4>' . $user->name . '</h4>';
    echo '<hr><pre>';
    print_r($user);
    echo '</pre>';
} else {
    echo '<h3>Not logged in</h3>';
    echo '<p><a href="?action=login">Log In</a></p>';
}

function apiRequest($url, $post=FALSE, $headers=array()) {
    $ch = curl_init($url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
    if ($post) {
        curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($post));
    }
    $headers[] = 'Accept: application/json';
    if ($_SESSION['access_token']) {
        $headers[] = 'Authorization: Bearer ' . $_SESSION['access_token'];
    }
    curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
    curl_setopt($ch,CURLOPT_USERAGENT,$_SERVER['HTTP_USER_AGENT']);
    $response = curl_exec($ch);
    return json_decode($response);
}
