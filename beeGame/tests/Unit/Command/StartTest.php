<?php

namespace Test\Unit\Command;

use ClientInterface\Cli;
use Command\Start;
use GamePlay\Game;
use State\InProgress as StateInProgress;

class StartTest extends \PHPUnit_Framework_TestCase
{
    public function testExecute()
    {
        $game = $this->getMock(Game::class, ['setState'], [new Cli()]);
        $game
            ->expects(static::once())
            ->method('setState')
            ->with(static::equalTo(new StateInProgress()))
            ->will(static::returnValue('OK'))
        ;
        $command = new Start();
        /** @var Game $game */
        $command->execute($game);
        // No need assert something, if method will not be invoked we'll receive fail.
    }
}
