<?php

$text = 'This is a Simple text.';

var_export([
    strpbrk($text, 'i'),
    strpbrk($text, 'S'),
    strpbrk($text, 'm'),
]);

/*
array (
  0 => 'is is a Simple text.',
  1 => 'Simple text.',
  2 => 'mple text.',
)
*/
