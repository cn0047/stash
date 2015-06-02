<?php

class Error
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

class MyError extends Error
{
    private function protectedGet()
    {
    }
}

/*
PHP Fatal error:  Access level to MyError::protectedGet() must be protected (as in class Error) or weaker
*/
