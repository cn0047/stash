<?php

require_once __DIR__ . '/vendor/autoload.php';

$text = new One\Text();
var_dump($text->get('test'));

/*
string(17) "One\Text::gettest"
*/
