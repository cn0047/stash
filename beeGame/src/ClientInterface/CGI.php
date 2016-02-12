<?php

namespace ClientInterface;

use State\StateInterface;

/**
 * CGI.
 *
 * This class provides ability to play game through CGI (rather HTTP).
 */
class CGI implements ClientInterfaceInterface
{
    /**
     * Constructor.
     *
     * Temporary solution.
     *
     * @throws \DomainException When someone try use this class.
     */
    public function __construct()
    {
        throw new \DomainException('CGI not implemented yet, please use cli.');
    }

    /**
     * {@inheritdoc}
     */
    public function getCommand(StateInterface $state)
    {
    }

    /**
     * {@inheritdoc}
     */
    public function outputStatistics(array $statistics)
    {
    }
}
