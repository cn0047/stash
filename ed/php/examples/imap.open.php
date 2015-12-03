<?php

$box = "{imap.gmail.com:993/imap/ssl}INBOX";
$user = "google-user@gmail.com";
$pass = "password";
// @see https://www.google.com/settings/security/lesssecureapps
$mbox = imap_open($box, $user, $pass);
if ($mbox === false) {
    var_export(imap_alerts());
    var_export(imap_last_error());
} else {
    $totalInTheBox = imap_num_msg($mbox);
    var_export($totalInTheBox);
}
