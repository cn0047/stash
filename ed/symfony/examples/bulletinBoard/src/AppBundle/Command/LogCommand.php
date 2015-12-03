<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * @example php app/console app:log --method=info pingLog
 * @see ./app/dev.my_cli.log
 */
class LogCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:log')
            ->setDescription('Log to monolog')
            ->addArgument(
                'message',
                InputArgument::REQUIRED,
                'What do you want to log?'
            )
            ->addOption(
                'method',
                null,
                InputOption::VALUE_OPTIONAL,
                'Method to log message',
                'info'
            )
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $message= $input->getArgument('message');
        $method = $input->getOption('method');
        $logger = $this->getContainer()->get('monolog.logger.my_cli');
        $logger->$method($message);
    }
}