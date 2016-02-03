<?php

namespace Command;

use GamePlay\Game;

interface CommandInterface
{
    public function execute(Game $game);
}
