<?php

final interface Queue
{
    public function add($item);
}

/*
PHP Parse error:  syntax error, unexpected 'interface' (T_INTERFACE), expecting class (T_CLASS)
*/
