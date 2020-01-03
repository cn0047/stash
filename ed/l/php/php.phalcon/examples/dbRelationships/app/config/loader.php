<?php

$loader = new \Phalcon\Loader();
$loader->registerDirs(
    array(
//        APP_PATH . $config->application->modelsDir,
        __DIR__ . '/../models',
        __DIR__ . '/../tasks'
    )
);
$loader->register();
