<?php

namespace Command;

use GamePlay\Game;
use State\End as StateEnd;
use State\InProgress as StateInProgress;

/**
 * Hit command.
 *
 * This command hit bee (main command of game). Exactly this - is a gist of this game.
 * This command can bring game to different states accordingly to fact that Queen bee is alive.
 */
class Hit implements CommandInterface
{
    /**
     * {@inheritdoc}
     */
    public function execute(Game $game)
    {
        $beeGang = $game->getBeeGang();
        $beeGang->randomHit();
        if ($beeGang->getIsQueenAlive()) {
            $game->setState(new StateInProgress());
        } else {
            $game->setState(new StateEnd());
        }
    }
}
