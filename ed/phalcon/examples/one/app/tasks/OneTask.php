<?php

/**
 * @example php app/cli.php one test foo bar
 */
class OneTask extends \Phalcon\Cli\Task
{
    /**
     * @param array $params
     */
    public function testAction(array $params)
    {
        echo sprintf('hello %s! Best regards %s from Task ONE.' . PHP_EOL, $params[0], $params[1]);
    }
}
