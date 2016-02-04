<?php

namespace Test\Unit\State;

use Command\CommandInterface;
use State\InProgress;

class InProgressTest extends \PHPUnit_Framework_TestCase
{
    /** @var InProgress */
    private $state;

    protected function setUp()
    {
        $this->state = new InProgress();
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
