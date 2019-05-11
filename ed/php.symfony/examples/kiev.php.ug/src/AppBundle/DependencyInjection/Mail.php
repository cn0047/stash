<?php

namespace AppBundle\DependencyInjection;

class Mail
{
    private $from = '';
    private $trackingService;
    private $twig;
    private $mailer;

    public function __construct($from, $trackingService, $twig, $mailer)
    {
        $this->from = $from;
        $this->trackingService = $trackingService;
        $this->twig = $twig;
        $this->mailer = $mailer;
    }

    public function mail($to, $subject, $text)
    {
        $headers  = "MIME-Version: 1.0\r\n";
        // $headers .= "Content-type: text/html; charset=iso-8859-1\r\n";
        $headers .= "From: {$this->from}\r\n";
        $result = mail($to, $subject, $text, $headers);
        $result = var_export($result, true);
        $this->trackingService->ioTrack('Mail.Send');
        $this->trackingService->log(
            "Send mail: $subject; to: $to; result: $result"
        );
        return $result;
    }

    public function swiftmail($to, $subject, $text)
    {
        $m = \Swift_Message::newInstance()
            ->setSubject($subject)
            ->setFrom($this->from)
            ->setTo($to)
            ->setBody($text)
        ;
        return $this->mailer->send($m);
    }

    public function sendMail($type, array $args)
    {
        return $this->$type($args);
    }

    final private function testMail(array $args)
    {
        $text = $this->twig->render('AppBundle:Mail:Test.txt.twig', $args);
        $subject = 'Test mail at '.date('Y-m-d H:i:s');
        return $this->mail($args['mail'], $subject, $text);
    }

    final private function testMail2(array $args)
    {
        $subject = 'Test mail at '.date('Y-m-d H:i:s');
        $text = $subject;
        return $this->swiftmail($args['mail'], $subject, $text);
    }

    final private function forgotPassword(array $args)
    {
        $text = $this->twig->render('AppBundle:Mail:ForgotPassword.txt.twig', $args);
        $subject = $this->get('translator')->trans('kpug.success.password_updated_2');
        return $this->mail($args['email'], $subject, $text);
    }
}
