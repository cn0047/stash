<?php

trait T1
{
    public function t1()
    {
        echo 200;
    }
}

trait T2
{
    use T1;

    public function t2()
    {
    }
}

class Bar
{
    use T2;
}

(new Bar())->t1();

/*
200
*/
