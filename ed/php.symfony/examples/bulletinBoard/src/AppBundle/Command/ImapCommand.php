<?php

namespace AppBundle\Command;

use Symfony\Bundle\FrameworkBundle\Command\ContainerAwareCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

/**
 * @example php app/console app:imap
 */
class ImapCommand extends ContainerAwareCommand
{
    protected function configure()
    {
        $this
            ->setName('app:imap')
            ->setDescription('Tweak imap')
        ;
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        $server = new \Fetch\Server('imap.gmail.com', 993);
        $user = "user@gmail.com";
        $pass = "pass";
        $server->setAuthentication($user, $pass);
        var_export(
            $server->listMailBoxes()
        );
        $messages = $server->getMessages();
        foreach ($messages as $message) {
            echo "Subject: {$message->getSubject()} \n";
        }
    }
}