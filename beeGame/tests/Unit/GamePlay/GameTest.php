<?php

namespace Test\Unit\GamePlay;

use ClientInterface\Cli;
use GamePlay\Game;
use State\Begin as StateBegin;
use Bee\Gang;

class StartTest extends \PHPUnit_Framework_TestCase
{
    /**
     * @todo Finalize this test.
     */
    public function testPlay()
    {
        // If be honest, i have no idea how i can test it
        // yet...
    }

    public function testGetLevel()
    {
        $game = new Game(new Cli());
        static::assertSame('LevelOne', $game->getLevel()->get());
    }


    public function testSetState()
    {
        $game = new Game(new Cli());
        $state = new StateBegin();
        $game->setState($state);
        // One available way to break encapsulation.
        $reflection = new \ReflectionObject($game);
        $property = $reflection->getProperty('state');
        $property->setAccessible(true);
        static::assertSame($state, $property->getValue($game));
    }

    public function testSetAndGetBeeGang()
    {
        $game = new Game(new Cli());
        $beeGang = new Gang();
        $game->setBeeGang($beeGang);
        static::assertSame($beeGang, $game->getBeeGang());
    }
}
