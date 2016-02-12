<?php

namespace Command;

use GamePlay\Game;
use State\End as StateEnd;

/**
 * Escape command.
 *
 * Command that allow to user leave the game.
 */
class Escape implements CommandInterface
{
    /**
     * {@inheritdoc}
     */
    public function execute(Game $game)
    {
        $game->setState(new StateEnd());
    }
}
