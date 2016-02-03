<?php

namespace Test\Unit\Command;

use Command\Escape;
use ClientInterface\Cli;
use GamePlay\Game;
use State\End as StateEnd;

class HitTest extends \PHPUnit_Framework_TestCase
{
    public function testExecute()
    {
        $game = $this->getMock(Game::class, ['setState'], [new Cli()]);
        $game
            ->expects(static::once())
            ->method('setState')
            ->with(static::equalTo(new StateEnd()))
            ->will(static::returnValue('OK'))
        ;
        $command = new Escape();
        $command->execute($game);
        // No need assert something, if method will not be invoked we'll receive fail.
    }
}
