<?php

namespace Test\Unit\ClientInterface;

use ClientInterface\Cli;

class CliTest extends \PHPUnit_Framework_TestCase
{
    /**
     * @todo Finalize this test.
     */
    public function testGetCommand()
    {
        // If be honest, i have no idea how i can test it
        // yet...
        // Moreover it's not exactly unit test,
        // here should be integration test...
    }

    public function testOutputStatistics()
    {
        $cli = new Cli();
        ob_start();
        $cli->outputStatistics(['Drone' => [50, 38]]);
        $output = ob_get_clean();
        static::assertContains('Drone: 50 38', $output);
    }
}
