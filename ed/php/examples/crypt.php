<?php

$hashed_password = crypt('mypassword');
if (hash_equals($hashed_password, crypt($user_input, $hashed_password))) {
    echo "Password verified!";
}
