<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * Read logs command.
 */
class ReadLogsCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:read-logs')
            ->setDescription('Read logs command')
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $logReader = $this->getContainer()->get('log_reader');
        $output->writeln($logReader->readLogs());
    }
}
