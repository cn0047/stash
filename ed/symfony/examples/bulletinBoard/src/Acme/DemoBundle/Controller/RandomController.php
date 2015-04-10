<?php

namespace Acme\DemoBundle\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;

class RandomController extends Controller
{
    public function indexAction($limit)
    {
        $number = rand(1, $limit);
        return new Response("Number: $number");
        // return $this->render('AcmeDemoBundle:Random:index.html.twig', array('number' => $number));
    }
}
