<?php

$box = "{imap.gmail.com:993/imap/ssl}INBOX";
$user = "codenamek2010@gmail.com";
$pass = "codenamek2010-007";
// @see https://www.google.com/settings/security/lesssecureapps
$mbox = imap_open($box, $user, $pass);
if ($mbox === false) {
    var_export([
        imap_alerts(),
        imap_last_error(),
    ]);
} else {
    $totalInTheBox = imap_num_msg($mbox);
    echo "In inbox: $totalInTheBox \n";
    if ($totalInTheBox > 0) {
        $id = 1;
        $message = imap_fetchbody($mbox, imap_uid($mbox, $id), 1.2, FT_UID | FT_PEEK);
        $message = quoted_printable_decode($message);
        $message = mb_convert_encoding($message, 'UTF-8', 'KOI8-R');
        var_export($message);
    }
}
