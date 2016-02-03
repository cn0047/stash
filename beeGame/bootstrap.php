<?php

error_reporting(E_ALL);

set_error_handler(
    function ($code, $description) {
        throw new ErrorException($description, $code);
    }
);

require_once './autoload.php';
