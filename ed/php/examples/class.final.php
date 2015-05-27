<?php

final class Father
{
    public function getSurname()
    {
    }
}

class Son extends Father
{
}

new Son;

/*
PHP Fatal error:  Class Son may not inherit from final class (Father)
*/
