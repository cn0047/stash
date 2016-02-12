<?php

namespace ClientInterface;

use State\StateInterface;
use Command\CommandInterface;

/**
 * Interface of interaction between client (gamer) and game.
 *
 * This interface describes rules how game interact with client (gamer) INPUT/OUTPUT layer.
 */
interface ClientInterfaceInterface
{
    /**
     * By interacting with client (gamer) understand which command should be executed next.
     *
     * @param StateInterface $state State of game.
     *
     * @return CommandInterface Kernel game command.
     */
    public function getCommand(StateInterface $state);

    /**
     * Output statistics about execution of command to client (gamer).
     *
     * @param array $statistics Statistics.
     */
    public function outputStatistics(array $statistics);
}
