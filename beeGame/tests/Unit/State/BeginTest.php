<?php

namespace Test\Unit\State;

use Command\CommandInterface;
use State\Begin;

class BeginTest extends \PHPUnit_Framework_TestCase
{
    /** @var Begin */
    private $state;

    protected function setUp()
    {
        $this->state = new Begin();
    }

    public function testGetPromptMessage()
    {
        static::assertNotEmpty($this->state->getPromptMessage());
    }

    public function testGetPromptedCommand()
    {
        static::assertInstanceOf(CommandInterface::class, $this->state->getPromptedCommand());
    }

    public function testGetNotPromptedCommand()
    {
        static::assertInstanceOf(CommandInterface::class, $this->state->getNotPromptedCommand());
    }
}
