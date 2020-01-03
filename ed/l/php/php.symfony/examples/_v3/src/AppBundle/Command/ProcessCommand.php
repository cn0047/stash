<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Process\Process;

class ProcessCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:process')
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        sleep(5);
        $process = new Process('ps aux | grep -v grep | grep app:process -c');
        $process->run();
        $count = $process->getOutput();
        $output->writeln(sprintf('Script works in %s threads.', trim($count)));
    }
}
