<?php

namespace GamePlay;

use Bee\Gang;
use ClientInterface\ClientInterfaceInterface;
use Command\CommandInterface;
use State\Begin as StateBegin;
use State\StateInterface;
use VO\NotEmptyString;

class Game
{
    /** @var NotEmptyString */
    private $level;

    /** @var StateInterface */
    private $state;

    /** @var ClientInterfaceInterface */
    private $interface;

    /** @var Gang */
    private $beeGang;

    public function __construct(ClientInterfaceInterface $clientInterface)
    {
        $this->interface = $clientInterface;
        // Initialize game.
        $this->level = new NotEmptyString('LevelOne');
        $this->setState(new StateBegin());
    }

    public function play()
    {
        /** @var CommandInterface $command */
        while ($command = $this->interface->getCommand($this->state)) {
            $command->execute($this);
            $this->interface->outputStatistics($this->beeGang->getStatistics());
        }
    }

    public function getLevel()
    {
        return $this->level;
    }

    public function setState(StateInterface $state)
    {
        $this->state = $state;
    }

    public function getBeeGang()
    {
        return $this->beeGang;
    }

    public function setBeeGang(Gang $beeGang)
    {
        $this->beeGang = $beeGang;
    }
}
