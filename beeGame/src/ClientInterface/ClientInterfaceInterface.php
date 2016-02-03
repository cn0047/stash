<?php

namespace ClientInterface;

use State\StateInterface;

interface ClientInterfaceInterface
{
    public function getCommand(StateInterface $state);

    public function outputStatistics(array $statistics);
}
