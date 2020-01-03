<?php

/**
 * It's ok for abstract class doesn't contains abstract methods.
 */
abstract class AppException extends Exception
{
}

class MyException extends AppException
{
}

$myException = new MyException();

var_export($myException);

/*
MyException::__set_state(array(
   'message' => '',
   'string' => '',
   'code' => 0,
   'file' => '/home/bond/web/kovpak/gh/ed/php/examples/class.abstract.extends.php',
   'line' => 11,
   'trace' =>
  array (
  ),
   'previous' => NULL,
))
*/
