<?php

namespace ClientInterface;

use State\StateInterface;

class Cli implements ClientInterfaceInterface
{
    public function getCommand(StateInterface $state)
    {
        echo "\n".$state->getPromptMessage();
        $confirmation = trim(fgets(STDIN));
        if ($confirmation === 'y') {
            return $state->getPromptedCommand();
        } else {
            return $state->getNotPromptedCommand();
        }
    }

    public function outputStatistics(array $statistics)
    {
        foreach ($statistics as $beeType => $points) {
            echo "\n$beeType: ".implode(' ', $points);
        }
    }
}
