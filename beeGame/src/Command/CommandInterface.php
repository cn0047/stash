<?php

namespace Command;

use GamePlay\Game;

/**
 * Game command interface.
 *
 * During playing game client (gamer) sends commands to the GamePlay,
 * exactly this commands will be executed.
 * Through commands we have interaction between client (gamer) and game (GamePlay).
 * Command - it's set of actions that move game from one state to another.
 * Each command must set new state of game.
 */
interface CommandInterface
{
    /**
     * Method that allow execute command.
     *
     * @param Game $game GamePlay.
     */
    public function execute(Game $game);
}
