<?php

namespace State;

use Command\Escape as CommandEscape;
use Command\Hit as CommandHit;

/**
 * In-progress game state.
 */
class InProgress implements StateInterface
{
    /**
     * {@inheritdoc}
     */
    public function getPromptMessage()
    {
        return 'Do you wish hit bee (y) or exit game (n)?';
    }

    /**
     * {@inheritdoc}
     */
    public function getPromptedCommand()
    {
        return new CommandHit();
    }

    /**
     * {@inheritdoc}
     */
    public function getNotPromptedCommand()
    {
        return new CommandEscape();
    }
}
