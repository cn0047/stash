<?php

class Script
{
    private function invokePrivate()
    {
        echo 'Invoked private.'.PHP_EOL;
    }

    protected function invokeProtected()
    {
        echo 'Invoked protected.'.PHP_EOL;
    }

    public function invokePublic()
    {
        echo 'Invoked public.'.PHP_EOL;
    }

    public function __call($name, $arguments)
    {
        echo "Magic invoke: $name.".PHP_EOL;
    }
}

$o = new Script;
$o->invokePrivate();
$o->invokeProtected();
$o->invokePublic();
$o->run();

/*
Magic invoke: invokePrivate.
Magic invoke: invokeProtected.
Invoked public.
Magic invoke: run.
*/
