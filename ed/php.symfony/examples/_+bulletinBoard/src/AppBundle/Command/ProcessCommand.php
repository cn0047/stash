<?php

namespace AppBundle\Command;

use duncan3dc\Helpers\Fork;
use Spork\ProcessManager;
use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * @example php app/console app:process --engine=spork --threadsCount=5
 */
class ProcessCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:process')
            ->addOption(
                'engine',
                null,
                InputOption::VALUE_OPTIONAL,
                'engine',
                'duncan3dc'
            )
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
        $engine = $input->getOption('engine');
        $threadsCount = $input->getOption('threadsCount');
        $this->$engine($threadsCount);
    }

    final private function duncan3dc($threadsCount)
    {
        $fork = new Fork();
        foreach (range(1, $threadsCount) as $i) {
            $fork->call(function () use ($i) {
                $r = $this->doSomeStuff($i);
                echo "duncan3dc: $r".PHP_EOL;
            });
        }
        $fork->wait();
    }

    final private function spork($threadsCount)
    {
        $manager = new ProcessManager();
        $that = $this;
        foreach (range(1, $threadsCount) as $i) {
            $manager->fork(function () use ($i, $that) {
                return $that->doSomeStuff($i);
            })->then(function ($fork) {
                echo "spork: Pid: {$fork->getPid()} says: {$fork->getResult()}".PHP_EOL;
            });
        }
    }

    final private function doSomeStuff($i)
    {
        sleep(2);
        return "Processed $i";
    }
}
