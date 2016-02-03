<?php

namespace State;

use Command\Escape as CommandEscape;
use Command\Hit as CommandHit;

class InProgress implements StateInterface
{
    public function getPromptMessage()
    {
        return 'Do you wish hit bee (y) or exit game (n)?';
    }

    public function getPromptedCommand()
    {
        return new CommandHit();
    }

    public function getNotPromptedCommand()
    {
        return new CommandEscape();
    }
}
