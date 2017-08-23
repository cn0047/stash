<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Filesystem\LockHandler;
use Symfony\Component\Process\Process;

class LockCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:lock')
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $t = $this->getContainer()->get('translator');
        $lock = new LockHandler(__FILE__);
        if ($lock->lock()) {
            $output->writeln($t->trans('Locked.'));
            sleep(5);
            $lock->release();
            $output->writeln($t->trans('Released.'));
        } else {
            $output->writeln($t->trans("Can't lock."));
        }
    }
}
