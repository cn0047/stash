<?php
/**
 * Chain of call.
 *
 * Like: $object->$function1()->$function2();
 */

class User
{
    public function getUser()
    {
        return new Admin;
    }
}

class Admin
{
    public function getName()
    {
        return 'Admin';
    }
}


$result = array_reduce(['getUser', 'getName'], function ($obj, $method) {
    return $obj->$method();
}, new User);
var_dump($result);
/*
string(5) "Admin"
*/
