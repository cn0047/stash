<?php

namespace AppBundle\Command;

use duncan3dc\Helpers\Fork;
use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * @example php app/console app:process --threadsCount=5
 */
class ProcessCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:process')
            ->addOption(
                'threadsCount',
                null,
                InputOption::VALUE_OPTIONAL,
                'count of threads',
                1
            )
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $threadsCount = $input->getOption('threadsCount');
        $fork = new Fork();
        foreach (range(1, $threadsCount) as $i) {
            $fork->call(function () use ($i) {
                $this->doSomeStuff($i);
            });
        }
        $fork->wait();
    }

    public function doSomeStuff($i)
    {
        echo "Process $i".PHP_EOL;
        sleep(2);
    }
}
