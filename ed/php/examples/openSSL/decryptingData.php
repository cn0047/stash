<?php

// Get the private Key
if (!$privateKey = openssl_pkey_get_private(file_get_contents('./private.key'))) {
    die('Private Key failed');
}
$a_key = openssl_pkey_get_details($privateKey);

// Decrypt the data in the small chunks
$chunkSize = ceil($a_key['bits']/8);
$output = '';

$encrypted = file_get_contents('/tmp/openSSL.tmp');

while ($encrypted) {
    $chunk = substr($encrypted, 0, $chunkSize);
    $encrypted = substr($encrypted, $chunkSize);
    $decrypted = '';
    if (!openssl_private_decrypt($chunk, $decrypted, $privateKey)) {
        die('Failed to decrypt data');
    }
    $output .= $decrypted;
}
openssl_free_key($privateKey);

// Uncompress the unencrypted data.
$output = gzuncompress($output);

echo 'Unencrypted Data: '.$output.PHP_EOL;
