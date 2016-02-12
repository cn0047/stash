<?php

namespace GamePlay;

use Bee\Gang;
use ClientInterface\ClientInterfaceInterface;
use Command\CommandInterface;
use State\Begin as StateBegin;
use State\StateInterface;
use VO\NotEmptyString;

/**
 * Game.
 *
 * This class encapsulates all game logic, and have access to all game objects.
 * Through commands game goes from one state to another, and maintain state of all necessary objects.
 */
class Game
{
    /** @var NotEmptyString Particular game level. */
    private $level;

    /** @var StateInterface State of game. */
    private $state;

    /** @var ClientInterfaceInterface Interface of interaction between client (gamer) and game. */
    private $interface;

    /** @var Gang Aggregate of bees. */
    private $beeGang;

    /**
     * Game constructor.
     *
     * Initialize game.
     *
     * @param ClientInterfaceInterface $clientInterface Client interface.
     *
     * @throws \InvalidArgumentException In case when level is not valid string.
     */
    public function __construct(ClientInterfaceInterface $clientInterface)
    {
        $this->interface = $clientInterface;
        // Initialize game.
        $this->level = new NotEmptyString('LevelOne');
        $this->setState(new StateBegin());
    }

    /**
     * This method provides game interaction.
     */
    public function play()
    {
        /** @var CommandInterface $command */
        while ($command = $this->interface->getCommand($this->state)) {
            $command->execute($this);
            // In case when beeGang initialized we can receive some statistics.
            if ($this->beeGang !== null) {
                $this->interface->outputStatistics($this->beeGang->getStatistics());
            }
        }
    }

    /**
     * Gets current level.
     *
     * @return NotEmptyString Level.
     */
    public function getLevel()
    {
        return $this->level;
    }

    /**
     * Sets new state of game.
     *
     * By this method commands can move game to the new state.
     *
     * @param StateInterface $state State of game.
     */
    public function setState(StateInterface $state)
    {
        $this->state = $state;
    }

    /**
     * Gets aggregate of bees.
     *
     * @return Gang Bee gang.
     */
    public function getBeeGang()
    {
        return $this->beeGang;
    }

    /**
     * Sets updated aggregate of bees.
     *
     * This method provides to commands ability to impact on bee aggregate.
     *
     * @param Gang $beeGang Bee gang.
     */
    public function setBeeGang(Gang $beeGang)
    {
        $this->beeGang = $beeGang;
    }
}
