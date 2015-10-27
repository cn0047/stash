<?php

define('OAUTH2_CLIENT_ID', '');
define('OAUTH2_CLIENT_SECRET', '');
define('OAUTH2_REDIRECT_URI', 'http://kiev-pug-dev.890m.com/oauth.gitHub.php');

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
        'redirect_uri' => OAUTH2_REDIRECT_URI,
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
        'redirect_uri' => OAUTH2_REDIRECT_URI,
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

/*
stdClass Object
(
    [login] => cn007b
    [id] => 5052324
    [avatar_url] => https://avatars.githubusercontent.com/u/5052324?v=3
    [gravatar_id] => 
    [url] => https://api.github.com/users/cn007b
    [html_url] => https://github.com/cn007b
    [followers_url] => https://api.github.com/users/cn007b/followers
    [following_url] => https://api.github.com/users/cn007b/following{/other_user}
    [gists_url] => https://api.github.com/users/cn007b/gists{/gist_id}
    [starred_url] => https://api.github.com/users/cn007b/starred{/owner}{/repo}
    [subscriptions_url] => https://api.github.com/users/cn007b/subscriptions
    [organizations_url] => https://api.github.com/users/cn007b/orgs
    [repos_url] => https://api.github.com/users/cn007b/repos
    [events_url] => https://api.github.com/users/cn007b/events{/privacy}
    [received_events_url] => https://api.github.com/users/cn007b/received_events
    [type] => User
    [site_admin] => 
    [name] => Vladimir Kovpak
    [company] => 
    [blog] => http://cn007b.tumblr.com/
    [location] => Kiev, Ukraine
    [email] => cn007b@gmail.com
    [hireable] => 1
    [bio] => 
    [public_repos] => 7
    [public_gists] => 0
    [followers] => 20
    [following] => 120
    [created_at] => 2013-07-20T07:02:25Z
    [updated_at] => 2015-10-26T10:08:37Z
    [private_gists] => 0
    [total_private_repos] => 0
    [owned_private_repos] => 0
    [disk_usage] => 4258
    [collaborators] => 0
    [plan] => stdClass Object
        (
            [name] => free
            [space] => 976562499
            [collaborators] => 0
            [private_repos] => 0
        )
)
*/

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
