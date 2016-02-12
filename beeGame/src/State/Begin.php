<?php

namespace State;

use \Command\Escape as CommandEscape;
use \Command\Start as CommandStart;

/**
 * Begin game state.
 */
class Begin implements StateInterface
{
    /**
     * {@inheritdoc}
     */
    public function getPromptMessage()
    {
        return 'Are you ready to start game (y/n)?';
    }

    /**
     * {@inheritdoc}
     */
    public function getPromptedCommand()
    {
        return new CommandStart();
    }

    /**
     * {@inheritdoc}
     */
    public function getNotPromptedCommand()
    {
        return new CommandEscape();
    }
}
