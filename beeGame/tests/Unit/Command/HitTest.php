<?php

namespace Test\Unit\Command;

use Bee\Gang;
use ClientInterface\Cli;
use Command\Hit;
use GamePlay\Game;
use State\End as StateEnd;
use State\InProgress as StateInProgress;

class HitTest extends \PHPUnit_Framework_TestCase
{
    public function testExecuteQueenAlive()
    {
        $game = $this->getMock(Game::class, ['setState'], [new Cli()]);
        $game
            ->expects(static::once())
            ->method('setState')
            ->with(static::equalTo(new StateInProgress()))
            ->will(static::returnValue('OK'))
        ;
        $beeGang = $this->getMock(Gang::class);
        $beeGang
            ->expects(static::once())
            ->method('randomHit')
            ->will(static::returnValue(null))
        ;
        $beeGang
            ->expects(static::once())
            ->method('getIsQueenAlive')
            ->will(static::returnValue(true))
        ;
        /** @var Game $game */
        /** @var Gang $beeGang */
        $game->setBeeGang($beeGang);
        $command = new Hit();
        $command->execute($game);
        // No need assert something, if method will not be invoked we'll receive fail.
    }

    public function testExecuteNotQueenAlive()
    {
        $game = $this->getMock(Game::class, ['setState'], [new Cli()]);
        $game
            ->expects(static::once())
            ->method('setState')
            ->with(static::equalTo(new StateEnd()))
            ->will(static::returnValue('OK'))
        ;
        $beeGang = $this->getMock(Gang::class);
        $beeGang
            ->expects(static::once())
            ->method('randomHit')
            ->will(static::returnValue(null))
        ;
        $beeGang
            ->expects(static::once())
            ->method('getIsQueenAlive')
            ->will(static::returnValue(false))
        ;
        /** @var Game $game */
        /** @var Gang $beeGang */
        $game->setBeeGang($beeGang);
        $command = new Hit();
        $command->execute($game);
        // No need assert something, if method will not be invoked we'll receive fail.
    }
}
