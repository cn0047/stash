<?php

class ExError
{
    private function privateGet()
    {
    }

    protected function protectedGet()
    {
    }

    public function publicGet()
    {
    }
}

class MyError extends ExError
{
    private function protectedGet()
    {
    }
}

/*
PHP Fatal error:  Access level to MyError::protectedGet() must be protected (as in class Error) or weaker
(public - OK).
*/
