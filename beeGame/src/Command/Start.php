<?php

namespace Command;

use Builder\Main as Builder;
use GamePlay\Game;
use State\InProgress as StateInProgress;

class Start implements CommandInterface
{
    public function execute(Game $game)
    {
        $builder = new Builder($game->getLevel());
        $builder->buildLevel();
        $game->setBeeGang($builder->getBeeGang());
        $game->setState(new StateInProgress());
    }
}
