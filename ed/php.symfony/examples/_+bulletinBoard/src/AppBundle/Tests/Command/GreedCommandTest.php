<?php

use AppBundle\Command\GreetCommand;
use Symfony\Bundle\FrameworkBundle\Console\Application;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;
use Symfony\Component\Console\Tester\CommandTester;

class GreedCommandTest extends KernelTestCase
{
    public function testExecute()
    {
        $kernel = $this->createKernel();
        $kernel->boot();
        $application = new Application($kernel);
        $application->add(new GreetCommand());
        $command = $application->find('demo:greet');
        $commandTester = new CommandTester($command);
        $commandTester->execute(
            array(
                'name'    => 'James Bond',
                '--yell'  => true,
            )
        );
        $this->assertRegExp('/JAMES\s{1}BOND/', $commandTester->getDisplay());
    }
}