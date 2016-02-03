<?php

namespace State;

use Command\Start as CommandStart;

class End implements StateInterface
{
    public function getPromptMessage()
    {
        return 'Game over. Start new game (y/n)?';
    }

    public function getPromptedCommand()
    {
        return new CommandStart();
    }

    public function getNotPromptedCommand()
    {
    }
}

