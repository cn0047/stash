<?php

$privateKey = openssl_pkey_new(array(
    'private_key_bits' => 2048, // Size of Key.
    'private_key_type' => OPENSSL_KEYTYPE_RSA,
));

// Save the private key to private.key file. Never share this file with anyone.
openssl_pkey_export_to_file($privateKey, './private.key');

// Generate the public key for the private key
$a_key = openssl_pkey_get_details($privateKey);
// Save the public key in public.key file. Send this file to anyone who want to send you the encrypted data.
file_put_contents('./public.key', $a_key['key']);

// Free the private Key.
openssl_free_key($privateKey);
