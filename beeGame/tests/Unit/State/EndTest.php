<?php

namespace Test\Unit\State;

use Command\CommandInterface;
use State\End;

class EndTest extends \PHPUnit_Framework_TestCase
{
    /** @var End */
    private $state;

    protected function setUp()
    {
        $this->state = new End();
    }

    public function testGetPromptMessage()
    {
        static::assertInternalType('string', $this->state->getPromptMessage());
    }

    public function testGetPromptedCommand()
    {
        static::assertInstanceOf(CommandInterface::class, $this->state->getPromptedCommand());
    }

    public function testGetNotPromptedCommand()
    {
        static::assertNull(null, $this->state->getNotPromptedCommand());
    }
}
