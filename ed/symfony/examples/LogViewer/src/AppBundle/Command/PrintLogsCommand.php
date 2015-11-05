<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * Print logs command.
 */
class PrintLogsCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:print-logs')
            ->setDescription('Print logs command')
            ->addArgument('filters', InputArgument::OPTIONAL, 'Filters', '')
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $filters = $input->getArgument('filters');
        $filters = json_decode($filters, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \RuntimeException('Parameter "filters" contains invalid JSON.');
        }
        // Obtain data.
        $em = $this->getContainer()->get('doctrine')->getManager();
        $data = $em->getRepository('AppBundle:Log')->findByFilters($filters);
        // Render data to cli output.
        $table = $this->getHelper('table');
        foreach ($data as &$v) {
            if (is_a($v['dateTime'], 'DateTime')) {
                $v['dateTime'] = $v['dateTime']->format('Y-m-d H:i:s');
            }
        }
        $table
            ->setHeaders([
                'id',
                'owner',
                'host',
                'user',
                'dateTime',
                'firstRequestLine',
                'status',
                'size',
                'userAgent',
            ])
            ->setRows($data)
        ;
        $table->render($output);
    }
}
