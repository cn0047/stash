<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * @example php app/console x
 */
class XCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('x')
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
    }
}
