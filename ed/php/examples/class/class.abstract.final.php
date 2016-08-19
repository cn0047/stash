<?php

abstract class AppException
{
    final abstract protected function setValue();
}

/*
PHP Fatal error:  Cannot use the final modifier on an abstract class member
*/
