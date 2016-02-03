<?php

namespace State;

use \Command\Escape as CommandEscape;
use \Command\Start as CommandStart;

class Begin implements StateInterface
{
    public function getPromptMessage()
    {
        return 'Are you ready to start game (y/n)?';
    }

    public function getPromptedCommand()
    {
        return new CommandStart();
    }

    public function getNotPromptedCommand()
    {
        return new CommandEscape();
    }
}
