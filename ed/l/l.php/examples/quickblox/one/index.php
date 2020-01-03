<?php

set_error_handler(
    function ($errorCode, $errorDescription, $errorFile, $errorLine, array $errorContext = []) {
        throw new Exception($errorDescription, $errorCode);
    }
);

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../config.php';
require_once __DIR__ . '/ChatBetweenAdminAndTargetUserNotFound.php';
require_once __DIR__ . '/QuickBloxBridge.php';
require_once __DIR__ . '/Command.php';

$action = $argv[1];
(new Command())->$action(...$argv);
