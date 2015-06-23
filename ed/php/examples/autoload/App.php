<?php

class App
{
    private $environment;
    private $application;

    public function __construct()
    {
        if (php_sapi_name() == 'cli') {
            $this->environment = 'cli';
            $this->application = 'application\Console';
        }
        new $this->application();
    }
}