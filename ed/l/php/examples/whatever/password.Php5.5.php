<?php

/**
 * Simple example.
 */
$password = 'SECRET';

/**
 * Prepare password to store at DB.
 */
$passwordHash = password_hash($password, PASSWORD_DEFAULT);
// Now $passwordHash contains string with length 60 characters,
// like: "$2y$10$wIUYtGKKVF3P3YO1O.IrI.Zav6j6H2/tCKEe4U/Mn8ykGVIyOlwZ."
// and it can be stored at DB.

/**
 * Verify password.
 */
$isPasswordVerified = password_verify($password, $passwordHash);
// In our case $isPasswordVerified will contains true.
// It means that password verified.
