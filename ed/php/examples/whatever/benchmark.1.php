<?php

var_export($_SERVER['REQUEST_URI']);
$uri = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);
var_dump($uri);

var_export($_GET);
var_export(file_get_contents("php://input"));
