<?php

trait A
{
    public function smallTalk()
    {
        return 'a';
    }
    public function bigTalk()
    {
        return 'A';
    }
}

trait B
{
    public function smallTalk()
    {
        return 'b';
    }
    public function bigTalk()
    {
        return 'B';
    }
}

class Talker
{
    use A, B {
        B::smallTalk insteadof A;
        A::bigTalk insteadof B;
    }
}

$o = new Talker;
var_export([
    $o->smallTalk(),
    $o->bigTalk(),
]);

/*
array (
  0 => 'b',
  1 => 'A',
)
*/
