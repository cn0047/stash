<?php

namespace ClientInterface;

use State\StateInterface;

class CGI implements ClientInterfaceInterface
{
    public function __construct()
    {
        throw new \DomainException('CGI not implemented yet, please use cli.');
    }

    public function getCommand(StateInterface $state)
    {
    }

    public function outputStatistics(array $statistics)
    {
    }
}
