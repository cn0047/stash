<?php

class MyError
{
    public function getDefaultLoggerFunction()
    {
        $f = function () {
            print(PHP_EOL . PHP_EOL);
            var_export($this);
        };
        return $f;
    }

    public function log(callable $cb)
    {
        $cb();
    }
}

class Logger
{
    public function getLoggerFunction()
    {
        $f = function () {
            print(PHP_EOL . PHP_EOL);
            var_export($this);
        };
        return $f;
    }
}

$varExport = function () {
    print(PHP_EOL . PHP_EOL);
    var_export($this);
};

$m = new MyError();
$m->log((new Logger())->getLoggerFunction()); // Logger
$m->log($varExport); // NULL
$m->log($m->getDefaultLoggerFunction()); // MyError
