<?php

class a
{
    public function a($x = 1)
    {
        $this->myvar = $x;
    }
}

class b extends a
{
    public $myvar;

    public function b($x = 2)
    {
        $this->myvar = $x;
        parent::a();
    }
}

$obj = new b;
echo $obj->myvar;
/*
1
*/
