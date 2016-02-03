<?php

namespace Command;

use GamePlay\Game;
use State\End as StateEnd;
use State\InProgress as StateInProgress;

class Hit implements CommandInterface
{
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
