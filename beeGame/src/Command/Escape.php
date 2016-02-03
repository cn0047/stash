<?php

namespace Command;

use GamePlay\Game;
use State\End as StateEnd;

class Escape implements CommandInterface
{
    public function execute(Game $game)
    {
        $game->setState(new StateEnd());
    }
}
