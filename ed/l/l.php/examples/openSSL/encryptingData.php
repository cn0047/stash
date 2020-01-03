<?php

// Data to be sent
$plaintext = 'Bond, James Bond!)';
// Compress the data to be sent
$plaintext = gzcompress($plaintext);

// Get the public Key of the recipient
$publicKey = openssl_pkey_get_public(file_get_contents('./public.key'));
$a_key = openssl_pkey_get_details($publicKey);

// Encrypt the data in small chunks and then combine and send it.
$chunkSize = ceil($a_key['bits']/8)-11;
$output = '';

while ($plaintext) {
    $chunk = substr($plaintext, 0, $chunkSize);
    $plaintext = substr($plaintext, $chunkSize);
    $encrypted = '';
    if (!openssl_public_encrypt($chunk, $encrypted, $publicKey)) {
        die('Failed to encrypt data');
    }
    $output .= $encrypted;
}
openssl_free_key($publicKey);

// This is the final encrypted data to be sent to the recipient
$encrypted = $output;

file_put_contents('/tmp/openSSL.tmp', $encrypted);
