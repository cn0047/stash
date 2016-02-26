<?php

/**
 * @example php app/cli.php main test world universe
 */
class MainTask extends \Phalcon\Cli\Task
{
    public function mainAction()
    {
        echo "\nThis is the default task and the default action \n";

        $this->console->handle(
            array(
                'task'   => 'main',
                'action' => 'testAdd',
            )
        );
    }

    /**
     * @param array $params
     */
    public function testAction(array $params)
    {
        echo sprintf('hello %s', $params[0]) . PHP_EOL;
        echo sprintf('best regards, %s', $params[1]) . PHP_EOL;
    }

    public function testAddAction()
    {
        echo "\nI will get printed too!\n";
    }
}
