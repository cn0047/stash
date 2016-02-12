<?php

namespace State;

use Command\Start as CommandStart;

/**
 * Eng game state (game over).
 */
class End implements StateInterface
{
    /**
     * {@inheritdoc}
     */
    public function getPromptMessage()
    {
        return 'Game over. Start new game (y/n)?';
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
        return null;
    }
}
